CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=route256/checkout/cmd/app

all: format build test lint

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})


.PHONY: bin
bin:
	GOBIN=$(BINDIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(BINDIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(BINDIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(BINDIR) go install github.com/envoyproxy/protoc-gen-validate@latest

PROTOC=PATH="$$PATH:$(BINDIR)" protoc
generate: bin
	mkdir -p pkg/loms_v1 &&\
 	$(PROTOC) -I ../libs/proto/loms/v1 -I ../vendor-proto \
		--go_out=pkg/loms_v1 --go_opt=paths=source_relative \
		--go-grpc_out=pkg/loms_v1 --go-grpc_opt paths=source_relative \
		--validate_out="lang=go,paths=source_relative:pkg/loms_v1" \
		../libs/proto/loms/v1/loms.proto
	mkdir -p pkg/productsvc_v1 &&\
	$(PROTOC) -I ../libs/proto/productsvc/v1 -I ../vendor-proto \
		--go_out=pkg/productsvc_v1 --go_opt=paths=source_relative \
		--go-grpc_out=pkg/productsvc_v1 --go-grpc_opt paths=source_relative \
		--validate_out="lang=go,paths=source_relative:pkg/productsvc_v1" \
		../libs/proto/productsvc/v1/productsvc.proto
	mkdir -p pkg/checkout_v1 &&\
	$(PROTOC) -I ../libs/proto/checkout/v1 -I ../vendor-proto \
		--go_out=pkg/checkout_v1 --go_opt=paths=source_relative \
		--go-grpc_out=pkg/checkout_v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out=pkg/checkout_v1 --grpc-gateway_opt paths=source_relative \
		--validate_out="lang=go,paths=source_relative:pkg/checkout_v1" \
		../libs/proto/checkout/v1/checkout.proto


.PHONY: migration-up
migration-up:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable" up

.PHONY: migration-down
migration-down:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable" down