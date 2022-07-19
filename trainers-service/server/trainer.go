package server

import (
	"context"
	"time"

	"github.com/patrickovm/go-microservices/trainers-service/model"
	"github.com/patrickovm/go-microservices/trainers-service/pb"
	"github.com/patrickovm/go-microservices/trainers-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trainer struct {
	repo repository.MongoRepo
	pb.UnimplementedTrainersServiceServer
}

func NewUser(repo repository.MongoRepo) *Trainer {
	return &Trainer{repo: repo}
}

func (t Trainer) Create(ctx context.Context, req *pb.CreateTrainerRequest) (*pb.CreateTrainerResponse, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	trainerId, err := t.repo.Create(c, &model.Trainer{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Age:           req.Age,
		FinishedGames: req.FinishedGames,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateTrainerResponse{Id: trainerId}, nil
}

func (t Trainer) Get(ctx context.Context, req *pb.GetTrainerRequest) (*pb.GetTrainerResponse, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	result, err := t.repo.FindOne(c, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetTrainerResponse{
		Id:            result.Id.Hex(),
		FirstName:     result.FirstName,
		LastName:      result.LastName,
		Age:           result.Age,
		FinishedGames: result.FinishedGames,
	}, nil
}

func (t Trainer) List(ctx context.Context, req *pb.ListTrainerRequest) (*pb.ListTrainerResponse, error) {
	data, err := t.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	var response pb.ListTrainerResponse
	for _, d := range *data {
		response.Trainers = append(response.Trainers, &pb.GetTrainerResponse{
			Id:            d.Id.Hex(),
			FirstName:     d.FirstName,
			LastName:      d.LastName,
			Age:           d.Age,
			FinishedGames: d.FinishedGames,
		})
	}

	return &response, nil
}

func (t Trainer) Update(ctx context.Context, req *pb.UpdateTrainerRequest) (*pb.UpdateTrainerResponse, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	trainerId, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	updatedTrainer, err := t.repo.Update(c, &model.Trainer{
		Id:            trainerId,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Age:           req.Age,
		FinishedGames: req.FinishedGames,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTrainerResponse{
		Id:            updatedTrainer.Id.Hex(),
		FirstName:     updatedTrainer.FirstName,
		LastName:      updatedTrainer.LastName,
		Age:           updatedTrainer.Age,
		FinishedGames: updatedTrainer.FinishedGames,
	}, nil
}
