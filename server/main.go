package main

import (
	"log"
	"net"

	"github.com/vyeve/grpc-tls/cert"
	"github.com/vyeve/grpc-tls/models"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", models.Address)
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

	log.Printf("start to listen on %s", models.Address[1:])
	log.Fatal(grpcServer.Serve(lis))
}
