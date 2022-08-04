package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/patrickovm/go-microservices/poke-service/pb"
	"github.com/patrickovm/go-microservices/poke-service/repository"
	"github.com/patrickovm/go-microservices/poke-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db, err := sql.Open("postgres", "postgres://example:example@localhost:5432/pokedb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostgresRepo(db)
	pokeServer := server.NewPokemon(repo)

	gs := grpc.NewServer()

	pb.RegisterPokeServiceServer(gs, pokeServer)

	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Print("Unable to create listener", "error", err)
		os.Exit(1)
	}

	log.Printf("Server started at port %v", 50052)

	gs.Serve(lis)
}
