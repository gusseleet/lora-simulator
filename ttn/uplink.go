package ttn
/*


import (
	"github.com/TheThingsNetwork/api/gateway"
	"github.com/TheThingsNetwork/api/protocol"
	"github.com/TheThingsNetwork/api/router"
	ttn_lorawan "github.com/TheThingsNetwork/api/protocol/lorawan"
	"lora_simulator/simulator"
)



func GenerateUplink(config *TTNConfig, generalConfig *simulator.GeneralConfig) (*router.UplinkMessage, error) {

	payload, err := simulator.GeneratePhyLoad(generalConfig, []byte{10, 10, 10, 10, 10})

	if err != nil {
		return nil, err
	}

	 return &router.UplinkMessage{
		GatewayMetadata: gateway.RxMetadata{
			GatewayID: config.gatewayID,
		},
		ProtocolMetadata:protocol.RxMetadata{
			Protocol: &protocol.RxMetadata_LoRaWAN{
				LoRaWAN: &ttn_lorawan.Metadata{
					CodingRate: config.codingRate,
					DataRate:   config.dataRate,
					Modulation: ttn_lorawan.Modulation_LORA,
				},
			},
		},

		Payload: payload,
	}, nil
}



*/
