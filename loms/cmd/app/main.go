package main

import (
	"context"
	_ "github.com/aitsvet/debugcharts"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os/signal"
	"route256/libs/mw/grpc/grpclogger"
	"route256/libs/mw/grpc/grpcvalidator"
	"route256/libs/postgresdb"
	"route256/libs/tx"
	"route256/loms/internal/config"
	grpchandler "route256/loms/internal/handler/grpc"
	"route256/loms/internal/jobs"
	"route256/loms/internal/repository"
	"route256/loms/internal/service"
	"route256/loms/pkg/loms_v1"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if err := config.InitConfig(); err != nil {
		log.Fatalf("[main] initilize config file: %v", err)
	}

	db, err := postgresdb.New(ctx, config.AppConfig.GetPostgresUrl())
	if err != nil {
		log.Fatalf("[main] failed to connect to postgres: %v", err)
	}
	repo := repository.NewRepository(db)
	_tx := tx.New(db)
	svc := service.New(repo, _tx)
	h := grpchandler.NewHandler(svc)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpclogger.Interceptor),
		grpc.ChainUnaryInterceptor(grpcvalidator.Interceptor),
	)
	reflection.Register(grpcServer)
	loms_v1.RegisterLOMSServer(grpcServer, h)

	errChan := make(chan error, 1)
	// preparation for gRPC server
	lis, err := net.Listen("tcp", config.AppConfig.GetGrpcServerAddr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()
	// http server for pprof and debugcharts
	// TODO: remove this, it is only for debugging
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// preparation for gRPC-Gateway
	conn, err := grpc.DialContext(
		context.Background(),
		lis.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()
	if err = loms_v1.RegisterLOMSHandler(context.Background(), mux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    config.AppConfig.GetHttpServerAddr(),
		Handler: mux,
	}
	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	// run jobs background
	j := jobs.NewJobs(svc)
	j.Start(ctx)
	defer j.Close()

	log.Println("starting server")
	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr)
	log.Printf("Serving gRPC on %s\n", lis.Addr().String())
	select {
	case err := <-errChan:
		log.Printf("[main] ListenAndServe got error: %v", err)
	case sig := <-ctx.Done():
		log.Printf("[main] terminated with: %v", sig)
	}
	// shutdown grpc server
	grpcServer.GracefulStop()

	// shutdown gateway server
	shoutDownTimeout := time.Millisecond * time.Duration(config.AppConfig.HttpServer.ShutdownTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), shoutDownTimeout)
	defer cancel()
	if err := gwServer.Shutdown(ctx); err != nil {
		log.Printf("[main] gwServer.Shutdown error: %v", err)
	}

	log.Println("server successfully shutdown")
}
