/**
Created by gel 2018-02-15
 */

package simulator

import (
	"lora_simulator/lora"
	"lora_simulator/logging"
	"github.com/sanity-io/litter"
)

func SendMessageLora(configLora * lora.LoraConfig, config *GeneralConfig, _payload string, lat int) error {

	a := []byte(_payload)

	client := lora.CreateConnection(configLora)
	payload, err := GeneratePhyLoad(config, a)
	logging.Log().Debug("Payload: ", litter.Sdump(payload))
	logging.Log().Debug("Payload (string): ", string(payload))


	if err != nil {
		return err
	}

	message := lora.TestMessage(payload, configLora.GatewayMac, lat)
	logging.Log().Debug("Message: ", litter.Sdump(message))

	pErr := lora.Publish(client, "gateway/"+configLora.GatewayMac+"/rx", message)

	logging.Log().Debug(message)


	return pErr

}