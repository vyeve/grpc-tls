package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vyeve/grpc-tls/cert"
	"github.com/vyeve/grpc-tls/models"

	"google.golang.org/grpc"
)

type client struct {
	client models.SSHClient
	logger *log.Logger
}

func main() {
	cl := initClient()
	for i := 0; i < 10; i++ {
		cl.makeRequest(i)
		time.Sleep(time.Second * 2)
	}
}

func initClient() *client {
	tlsCredentials, err := cert.LoadClientTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	conn, err := grpc.Dial(
		models.Address,
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &client{
		client: models.NewSSHClient(conn),
		logger: log.New(os.Stdout, "client: ", log.Ldate|log.Ltime|log.LUTC),
	}
}

func (c *client) makeRequest(n int) {
	resp, err := c.client.Echo(context.Background(), &models.Request{
		Id:   uint64(n),
		Body: fmt.Sprintf("client message #%d", n),
	})
	if err != nil {
		c.logger.Printf("ERROR: make request #%d. err: %v", n, err)
		return
	}
	c.logger.Printf("INFO: response server on n=%d: %s", n, resp.Body)
}
