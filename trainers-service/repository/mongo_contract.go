package repository

import (
	"context"

	"github.com/patrickovm/go-microservices/trainers-service/model"
)

type MongoRepo interface {
	Create(ctx context.Context, trainer *model.Trainer) (string, error)
	FindOne(ctx context.Context, id string) (*model.Trainer, error)
	List(ctx context.Context) (*model.Trainer, error)
	Update(ctx context.Context, trainer *model.Trainer) (*model.Trainer, error)
}
