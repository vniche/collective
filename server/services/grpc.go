package services

import (
	"log"
	"net"
	"os"
)

func StartGrpc() {
	service := New()

	grpcPort, ok := os.LookupEnv("GRPC_PORT")
	if !ok {
		grpcPort = "0.0.0.0:50051"
	}

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	defer service.Shutdown()

	log.Printf("server listening at %v", lis.Addr())
	if err := service.server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
