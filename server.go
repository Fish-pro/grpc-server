package main

import (
	"github.com/Fish-pro/grpc-server/helper"
	"github.com/Fish-pro/grpc-server/services"
	"google.golang.org/grpc"

	"log"
	"net"
	"os"
)

func main() {
	//cred, err := credentials.NewServerTLSFromFile("keys/server.crt","keys/server_no_passwd.key")
	//if err != nil{
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCred()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))   // 商品服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OrderService)) // 订单服务

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	rpcServer.Serve(lis)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request.Proto)
	//	fmt.Println(request.Header)
	//	fmt.Println(request)
	//	rpcServer.ServeHTTP(writer, request)
	//})
	//httpServer := &http.Server{
	//	Addr:    ":8080",
	//	Handler: mux,
	//}
	//
	//httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_passwd.key")

}
