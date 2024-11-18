package grpc

import (
	"context"
	"fmt"

	"github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/grpc/pb"
)

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
			Response: fmt.Sprintf("%v", err),
		}, nil
	}

	return &pb.CreateCharacterResponse{
		Response: "Created",
	}, nil
}

func (s *server) UpdateCharacter(_ context.Context, update *pb.UpdateCharacterRequest) (*pb.UpdateCharacterResponse, error) {
	return &pb.UpdateCharacterResponse{
		Response: "Updated",
	}, nil
}

func (s *server) DeleteCharacter() {

}

func (s *server) GetCharacter() {

}
