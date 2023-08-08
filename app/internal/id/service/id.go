package id

import (
	"context"
	pb "id-manager/internal/id/pb"
	"strconv"
)

type Server struct {
	pb.UnimplementedIdServiceServer
}

func (s *Server) Create(ctx context.Context, empty *pb.Empty) (*pb.CreateIdResponse, error) {
	id := GenerateId()

	return &pb.CreateIdResponse{Id: strconv.FormatUint(id, 10)}, nil
}
