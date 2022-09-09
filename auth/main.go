package main

import (
	"github.com/ElegantSoft/go-micro-example2/auth/handler"
	pb "github.com/ElegantSoft/go-micro-example2/auth/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "auth"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterAuthHandler(srv.Server(), new(handler.Auth))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
