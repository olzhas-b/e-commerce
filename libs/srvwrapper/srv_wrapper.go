package srvwrapper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

const (
	defaultContentType = "application/json"
)

type Wrapper[Req Validator, Resp any] struct {
	fn func(ctx context.Context, req Req) (Resp, error)
}

type Validator interface {
	Validate() error
}

func New[Req Validator, Resp any](fn func(ctx context.Context, req Req) (Resp, error)) *Wrapper[Req, Resp] {
	return &Wrapper[Req, Resp]{
		fn: fn,
	}
}

func (wr *Wrapper[Req, Resp]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("request begin")

	var req Req
	// decode request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("decode err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// validate if request implements Validator interface
	validator, ok := any(req).(Validator)
	if ok {
		err := validator.Validate()
		if err != nil {
			log.Printf("validator error: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// execute handler
	resp, err := wr.fn(context.Background(), req)
	if err != nil {
		log.Printf("handler error: %v\n", err)
		return
	}

	// encode response
	bs, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set content type
	w.Header().Add("accept", defaultContentType)
	w.Header().Add("Content-Type", defaultContentType)

	// write response
	_, err = w.Write(bs)
	if err != nil {
		log.Printf("write err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("request end")
}
