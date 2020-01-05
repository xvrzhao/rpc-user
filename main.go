package main

import (
	"fmt"
	"github.com/micro-stacks/rpc-user/db"
	pb "github.com/micro-stacks/rpc-user/proto"
	"github.com/micro-stacks/rpc-user/server"
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
