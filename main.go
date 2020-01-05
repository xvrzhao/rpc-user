package main

import (
	"code.aliyun.com/microstack/rpc-user/db"
	pb "code.aliyun.com/microstack/rpc-user/proto"
	"code.aliyun.com/microstack/rpc-user/server"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func migrate() {
	err := db.Migrate()
	if err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}
}

func serve() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, new(server.UserServer))

	fmt.Println("gRPC server is started.")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	migrate()
	serve()
}
