package repository

import (
	"database/sql"

	"github.com/patrickovm/go-microservices/poke-service/model"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (p PostgresRepo) Create(pokemon model.Pokemon) (int64, error) {
	err := p.db.QueryRow(
		"INSERT INTO pokemons(TrainerId, Name, HasBadge) VALUES($1, $2) RETURNING Id",
		pokemon.TrainerId, pokemon.Name, pokemon.HasBadge).Scan(&pokemon.Id)

	if err != nil {
		return 0, err
	}
	return pokemon.Id, nil
}

func (p PostgresRepo) Get(trainerId string) ([]model.Pokemon, error) {
	rows, err := p.db.Query("SELECT * FROM pokemons WHERE TrainerId=$1", trainerId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pokemons := []model.Pokemon{}

	for rows.Next() {
		var p model.Pokemon
		if err := rows.Scan(&p.Id, &p.TrainerId, &p.Name, &p.HasBadge); err != nil {
			return nil, err
		}
		pokemons = append(pokemons, p)
	}

	return pokemons, nil
}

func (p PostgresRepo) Update(id int64, pokemon model.Pokemon) (model.Pokemon, error) {
	_, err :=
		p.db.Exec("UPDATE pokemons SET TrainerId=$1, Name=$2, HasBadge=$3 WHERE id=$4",
			pokemon.Name, pokemon.HasBadge, pokemon.Id)

	if err != nil {
		return model.Pokemon{}, err
	}
	return pokemon, nil
}
