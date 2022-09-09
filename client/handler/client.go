package handler

import (
	"context"
	pb "github.com/ElegantSoft/go-micro-example2/client/proto"
)

type Client struct{}

func (c Client) Call(ctx context.Context, empty *pb.Empty, response *pb.UsernameResponse) error {
	response.Username = "ahmed"
	return nil
}

func NewClient() pb.ClientHandler {
	return &Client{}
}
