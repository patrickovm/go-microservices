package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/patrickovm/go-microservices/trainers-service/pb"
	"github.com/patrickovm/go-microservices/trainers-service/repository"
	"github.com/patrickovm/go-microservices/trainers-service/server"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	mongoUri := viper.GetString("MONGO_URI")
	mongoUsername := viper.GetString("MONGO_USER")
	mongoPassword := viper.GetString("MONGO_PASSWORD")
	db := viper.GetString("MONGO_DB")
	collection := viper.GetString("MONGO_COLLECTION")

	client, _ := repository.InitializeMongo(mongoUri, mongoUsername, mongoPassword)
	defer client.Disconnect(context.Background())

	mongoCollection := client.Database(db).Collection(collection)

	repo := repository.NewMongoRepository(mongoCollection)
	userServer := server.NewTrainer(repo)

	gs := grpc.NewServer()

	pb.RegisterTrainersServiceServer(gs, userServer)

	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("Unable to create listener", "error", err)
		os.Exit(1)
	}

	log.Printf("Server started at port %v", 50051)

	gs.Serve(lis)
}
