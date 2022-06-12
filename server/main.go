package main

import (
	"log"
	"net"

	"github.com/vyeve/grpc-ssh-server/cert"
	"github.com/vyeve/grpc-ssh-server/models"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50880")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tlsCredentials, err := cert.LoadServerTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	models.RegisterSSHServer(grpcServer, NewServer())

	log.Printf("start to listen on %s", "50880")
	log.Fatal(grpcServer.Serve(lis))
}
