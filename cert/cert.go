package cert

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"errors"

	"google.golang.org/grpc/credentials"
)

var (
	//go:embed ca-cert.pem
	serverCertPEM []byte
	//go:embed ca-key.pem
	serverKeyPEM []byte
	//go:embed ca-cert.pem
	caCert []byte
	//go:embed client-cert.pem
	clientCertPEM []byte
	//go:embed client-key.pem
	clientKeyPEM []byte
)

func LoadServerTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.X509KeyPair(serverCertPEM, serverKeyPEM)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func LoadClientTLSCredentials() (credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("failed to add server CA's certificate")
	}
	clientCert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		return nil, err
	}
	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
