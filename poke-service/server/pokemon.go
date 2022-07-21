package server

import (
	"context"

	"github.com/patrickovm/go-microservices/poke-service/model"
	"github.com/patrickovm/go-microservices/poke-service/pb"
	"github.com/patrickovm/go-microservices/poke-service/repository"
)

type Pokemon struct {
	repo repository.Repository
	pb.UnimplementedPokeServiceServer
}

func NewPokemon(repo repository.Repository) *Pokemon {
	return &Pokemon{repo: repo}
}

func (p Pokemon) Create(ctx context.Context, req *pb.CreatePokemonRequest) (*pb.CreatePokemonResponse, error) {

	id, err := p.repo.Create(model.Pokemon{
		TrainerId: req.TrainerId,
		Name:      req.Name,
		HasBadge:  req.HasBadge,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreatePokemonResponse{Id: id}, nil
}

func (p Pokemon) GetAllFromOneTrainer(ctx context.Context, req *pb.GetAllFromOneTrainerRequest) (*pb.GetAllFromOneTrainerResponse, error) {
	data, err := p.repo.Get(req.TrainerId)

	if err != nil {
		return nil, err
	}

	var response pb.GetAllFromOneTrainerResponse

	for _, d := range data {
		response.Pokelist = append(response.Pokelist, &pb.UpdatePokemonRequest{
			Id:        d.Id,
			TrainerId: d.TrainerId,
			Name:      d.Name,
			HasBadge:  d.HasBadge,
		})
	}

	return &response, nil
}

func (p Pokemon) Update(ctx context.Context, req *pb.UpdatePokemonRequest) (*pb.UpdatePokemonResponse, error) {
	updatedPokemon, err := p.repo.Update(req.Id, model.Pokemon{
		TrainerId: req.TrainerId,
		Name:      req.Name,
		HasBadge:  req.HasBadge,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePokemonResponse{
		Id:        updatedPokemon.Id,
		TrainerId: updatedPokemon.TrainerId,
		Name:      updatedPokemon.Name,
		HasBadge:  updatedPokemon.HasBadge,
	}, nil
}
