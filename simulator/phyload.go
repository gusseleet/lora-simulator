package simulator

import (
	"github.com/brocaar/lorawan"
)

func GeneratePhyLoad(config *GeneralConfig, payload []byte) ([]byte, error) {

	phy := lorawan.PHYPayload{
		MHDR: lorawan.MHDR{

			MType: lorawan.ConfirmedDataUp,
			Major: lorawan.LoRaWANR1,
		},
		MACPayload: &lorawan.MACPayload{
			FHDR: lorawan.FHDR {
				DevAddr: lorawan.DevAddr(config.DevAddr),
				FCnt:    1,
			},
			FPort:      &config.FPort,
			FRMPayload: []lorawan.Payload{&lorawan.DataPayload{Bytes: payload}},
		},
	}

	if err := phy.EncryptFRMPayload(config.AppSKey); err != nil {
		return nil, err
	}

	if err := phy.SetMIC(config.NwkSKey); err != nil {
		return nil, err
	}

	bytes, err := phy.MarshalText()

	if err != nil {
		return nil, err
	}


	return bytes, nil
}