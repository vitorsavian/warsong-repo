package grpc

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/grpc/pb"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/controllers"
)

type server struct {
	pb.UnimplementedCharacterServer
	pb.UnimplementedPingServer
	pb.UnimplementedWeaponServer
	characterController *controllers.CharacterController
}

func NewServer() *server {
	characterController, err := controllers.GetCharacterController()
	if err != nil {
		logrus.Fatalf("Failed to get character controller: %v", err)
	}

	return &server{
		characterController: characterController,
	}
}

func (s *server) PingServer(_ context.Context, ping *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Ping: 1,
	}, nil
}
