package lora

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"encoding/json"
	"lora_simulator/logging"
)
type LoraConfig struct {
	MqttAdress		string
	MqttPassword	string
	MqttUser		string
	GatewayMac		string
	AppEUI			[8]byte
	DevEUI			[8]byte
	NwsHexkey		string
	AppHexKey		string
	DevHexAddr		string
	Bandwidth		string
	Modulation 		string
	Frequency		int64
	CodeRate		string

}

func CreateConnection(config *LoraConfig) MQTT.Client {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(config.MqttAdress)
	//opts.SetUsername(config.MqttUser)
	//opts.SetPassword(config.MqttPassword)

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logging.Log().Error("Error with token: ", token.Error())
	}

	return client
}

func Publish(client MQTT.Client, topic string, v interface{}) error {
	bytes, err := json.Marshal(v)

	if err != nil {
		return err
	}

	if token := client.Publish(topic, 0, false, bytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}