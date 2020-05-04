package server

import (
	"log"
	"net"

	"github.com/amanchourasiya/juna/pkg/message"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func StartServer() {

	log.Printf("Starting server ...\n")

	userService := &UserService{}
	lis, err := net.Listen("tcp", "localhost"+port)
	if err != nil {
		log.Fatalf("Failed to start server %v\n", err)
	}
	//opt := []grpc.ServerOption{grpc.WithInsecure()}

	grpcServer := grpc.NewServer()
	message.RegisterUserServiceServer(grpcServer, userService)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve grpc service %v\n", err)
	}

	log.Printf("Server started successfully on port%s\n", port)
}
