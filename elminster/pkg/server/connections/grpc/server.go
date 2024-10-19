package grpc

import (
	"context"

	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/grpc/proto"
)

type Server struct {
	proto.UnimplementedCharacterServer
}

func (s *Server) CreateCharacter(_ context.Context, chracter *proto.CreateCharacterRequest) (*proto.CreateCharacterResponse, error) {
	return &proto.CreateCharacterResponse{
		Message: "Created",
	}, nil
}
