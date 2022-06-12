package main

import (
	"context"
	"log"
	"os"

	"github.com/vyeve/grpc-ssh-server/models"
)

type server struct {
	models.UnimplementedSSHServer
	logger *log.Logger
}

func NewServer() models.SSHServer {
	return &server{
		logger: log.New(os.Stdout, "server: ", log.Ldate|log.Ltime|log.LUTC),
	}
}

func (s *server) Echo(ctx context.Context, req *models.Request) (*models.Response, error) {
	s.logger.Printf("Request: #%d. Body: %s", req.Id, req.Body)
	return &models.Response{
		Id:   req.Id,
		Body: "echo server [" + req.Body + "]",
	}, nil
}
