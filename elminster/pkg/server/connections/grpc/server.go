package grpc

import (
	"context"

	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/grpc/pb"
)

type server struct {
	pb.UnimplementedCharacterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateCharacter(_ context.Context, chracter *pb.CreateCharacterRequest) (*pb.CreateCharacterResponse, error) {
	return &pb.CreateCharacterResponse{
		Message: "Created",
	}, nil
}
