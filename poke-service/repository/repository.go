package repository

import "github.com/patrickovm/go-microservices/poke-service/model"

type Repository interface {
	Create(pokemon model.Pokemon) (int64, error)
	Get(trainerId string) ([]model.Pokemon, error)
	Update(id int64, pokemon model.Pokemon) (model.Pokemon, error)
}
