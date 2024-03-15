package main

import (
	"github.com/pandishpancheta/listing-service/pkg/ai"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pandishpancheta/listing-service/pkg/config"
	"github.com/pandishpancheta/listing-service/pkg/db"
	"github.com/pandishpancheta/listing-service/pkg/pb"
	"github.com/pandishpancheta/listing-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.LoadConfig()

	db := db.Init(cfg)

	model := ai.InitModel()

	lis, err := net.Listen("tcp", "localhost:"+cfg.TCP_PORT)
	if err != nil {
		panic(err)
	}

	ls := services.NewListingsService(db, cfg, model)

	grpcServer := grpc.NewServer()

	pb.RegisterListingsServiceServer(grpcServer, ls)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
