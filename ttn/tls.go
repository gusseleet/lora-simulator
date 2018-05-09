package ttn
/*
package ttn

import (
	"io/ioutil"
	"crypto/x509"
	"crypto/tls"
	"errors"
	"google.golang.org/grpc/credentials"
)

func readFile(path string) ([]byte, error){
	return ioutil.ReadFile(path)
}

func createTLSConfig(path string, netHost string) (credentials.TransportCredentials, error){

	certPool := x509.NewCertPool()
	ca, err := readFile(path)

	if( err != nil) {
		return nil, err
	}

	if ok := certPool.AppendCertsFromPEM([]byte(ca)); !ok {
		return nil, errors.New("can't append cert")
	}

	return credentials.NewTLS(&tls.Config{ServerName:netHost, RootCAs:certPool}), nil

}
*/
