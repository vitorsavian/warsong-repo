package grpc

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/grpc/pb"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/controllers"
)

type server struct {
	pb.UnimplementedCharacterServer
	characterController *controllers.CharacterController
}

func NewServer() *server {
	characterController, err := controllers.GetController()
	if err != nil {
		logrus.Fatalf("Failed to get character controller: %v", err)
	}

	return &server{
		characterController: characterController,
	}
}

func (s *server) CreateCharacter(_ context.Context, character *pb.CreateCharacterRequest) (*pb.CreateCharacterResponse, error) {

	// Generate the struct that will be send to the controller
	body := &adapter.CharacterCreationRequestAdapter{
		Level: int8(character.Level),
		Name:  character.Name,
		Stats: adapter.StatsCharacterCreationRequestAdapter{
			Str: int8(character.Stats.Str),
			Dex: int8(character.Stats.Dex),
			Con: int8(character.Stats.Con),
			Cha: int8(character.Stats.Cha),
			Int: int8(character.Stats.Int),
			Wil: int8(character.Stats.Wil),
		},
	}

	if _, err := s.characterController.CreateCharacter(body); err != nil {
		return &pb.CreateCharacterResponse{
			Message: fmt.Sprintf("%v", err),
		}, nil
	}

	return &pb.CreateCharacterResponse{
		Message: "Created",
	}, nil
}
