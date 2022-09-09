package main

import (
	"github.com/ElegantSoft/go-micro-example2/client/handler"
	pb "github.com/ElegantSoft/go-micro-example2/client/proto"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "client"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()
	pb.RegisterClientHandler(srv.Server(), handler.NewClient())
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
