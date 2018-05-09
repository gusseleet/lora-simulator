package ttn
/*


import (
	rc "github.com/TheThingsNetwork/api/router/routerclient"
	"context"
	"google.golang.org/grpc"
)

type TTNConfig struct {
	address string
	certPath string
	netHost string
	accountServer string
	gatewayKey string
	gatewayID string
	buffersize int
	dataRate string
	codingRate string
}

func CreateGatewayStream(config *TTNConfig) (rc.GenericStream, error)  {

	tlsConfig, err := createTLSConfig(config.certPath, config.netHost)

	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(config.address, grpc.WithTransportCredentials(tlsConfig))

	if err != nil {
		return nil, err
	}

	token, err := getToken(config.accountServer, config.gatewayKey, config.gatewayID)

	if err != nil {
		return nil, err
	}

	routerClient := rc.NewClient(rc.ClientConfig{BackgroundContext:context.Background(), BufferSize: config.buffersize})
	routerClient.AddServer(config.address, conn)

	return routerClient.NewGatewayStreams(config.gatewayID, token, false), nil


}*/
