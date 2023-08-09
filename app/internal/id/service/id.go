package id

import (
	"context"
	"strconv"
	pb "vrsf-playground/id-manager/internal/id/pb"
)

type Server struct {
	pb.UnimplementedIdServiceServer
}

func (s *Server) Create(ctx context.Context, empty *pb.Empty) (*pb.CreateIdResponse, error) {
	id := GenerateId()

	return &pb.CreateIdResponse{Id: strconv.FormatUint(id, 10)}, nil
}
