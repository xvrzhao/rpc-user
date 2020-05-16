package main

import (
	"fmt"
	"github.com/micro-stacks/rpc-user/db/models"
	pb "github.com/micro-stacks/rpc-user/proto"
	"github.com/micro-stacks/rpc-user/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var listener net.Listener

func migrateDB() {
	err := models.Migrate()
	if err != nil {
		log.Fatalf("failed to migrate db: %s", err.Error())
	}
}

func createListener() {
	var err error
	if listener, err = net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT"))); err != nil {
		log.Fatalf("failed to create TCP listener: %s", err.Error())
	}
}

func init() {
	migrateDB()
	createListener()
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, new(server.UserServer))

	fmt.Println("gRPC server listening ...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
