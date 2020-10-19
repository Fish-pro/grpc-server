package main

import (
	"context"
	"github.com/Fish-pro/grpc-server/helper"
	"github.com/Fish-pro/grpc-server/services"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCred())}
	err := services.RegisterProdServiceHandlerFromEndpoint(
		context.Background(),
		gwmux,
		"localhost:8081",
		opt,
	)
	if err != nil {
		log.Println(">>>", err.Error())
		os.Exit(1)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()

}
