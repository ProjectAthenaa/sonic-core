package certificate_module

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"errors"
	"google.golang.org/grpc/credentials"
)

//go:embed cert/ca-cert.pem
var pemClientCA []byte

//go:embed cert/server-cert.pem
var certPEM []byte

//go:embed cert/server-key.pem
var keyPEM []byte

func LoadCertificate() (credentials.TransportCredentials, error) {

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, errors.New("failed to add client CA's certificate")
	}

	serverCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	return credentials.NewTLS(config), nil
}

//go:embed local_cert/ca-cert.pem
var caCert []byte

//go:embed local_cert/server-cert.pem
var certPEMLCL []byte

//go:embed local_cert/server-key.pem
var keyPEMLCL []byte

func LoadTestCertificate() (credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("failed to add client CA's certificate")
	}

	serverCert, err := tls.X509KeyPair(certPEMLCL, keyPEMLCL)
	if err != nil {
		return nil, err
	}

	//config := &tls.Config{
	//	Certificates: []tls.Certificate{serverCert},
	//	ClientAuth:   tls.RequireAndVerifyClientCert,
	//	ClientCAs:    certPool,
	//}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
		InsecureSkipVerify: true,
	}

	return credentials.NewTLS(config), nil
}

func LoadClientTestCertificate() (credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("failed to add client CA's certificate")
	}

	config := &tls.Config{
		RootCAs: certPool,
		InsecureSkipVerify: true,
	}
	return credentials.NewTLS(config), nil
}
