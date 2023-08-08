package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"id-manager/config"
	pb "id-manager/internal/id/pb"
	idGenerator "id-manager/internal/id/service"
)

func InitServer(c config.Config) {
	log.Printf("Starting gRPC on :%s ...", c.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.GrpcPort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterIdServiceServer(grpcServer, &idGenerator.Server{})
	healthpb.RegisterHealthServer(grpcServer, health.NewServer())

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
