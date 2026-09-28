package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/nitro53xp-coder/-arbforge-grpc-/gen"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterArbForgeServiceServer(server, &ArbForgeServer{})

	log.Println("gRPC server listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
