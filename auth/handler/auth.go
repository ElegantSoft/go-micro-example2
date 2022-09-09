package handler

import (
	"context"
	pb "github.com/ElegantSoft/go-micro-example2/auth/proto"
	"go-micro.dev/v4/logger"
)

type Auth struct{}

func (a Auth) Login(ctx context.Context, request *pb.LoginRequest, response *pb.LoginResponse) error {
	logger.Debug("will login")
	response.Token = request.Username + " hello man"
	return nil
}

func NewAuthService() pb.AuthHandler {
	return &Auth{}
}
