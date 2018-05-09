package main

import (
	"github.com/brocaar/lorawan"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"fmt"
	//"sync"
	"time"
	"github.com/brocaar/lorawan/band"
	//"github.com/loraserver/api/gw"
	"encoding/json"
	"strings"
	"lora_simulator/redis"
)


type Duration time.Duration
type EUI64 [8]byte

type RXInfo struct {
	MAC               lorawan.EUI64			`json:"mac"`
	Time              *time.Time    `json:"time,omitempty"`
	TimeSinceGPSEpoch *Duration     `json:"timeSinceGPSEpoch,omitempty"` // Time since GPS epoch (1980-01-06, only set when the gateway has a GPS time source)
	Timestamp         uint32        `json:"timestamp"`                   // gateway internal receive timestamp with microsecond precision, will rollover every ~ 72 minutes
	Frequency         int           `json:"frequency"`                   // frequency in Hz
	Channel           int           `json:"channel"`                     // concentrator IF channel used for RX
	RFChain           int           `json:"rfChain"`                     // RF chain used for RX
	CRCStatus         int           `json:"crcStatus"`                   // 1 = OK, -1 = fail, 0 = no CRC
	CodeRate          string        `json:"codeRate"`                    // ECC code rate
	RSSI              int           `json:"rssi"`                        // RSSI in dBm
	LoRaSNR           float64       `json:"loRaSNR"`                     // LoRa signal-to-noise ratio in dB
	Size              int           `json:"size"`                        // packet payload size
	DataRate          band.DataRate `json:"dataRate"`                    // RX datarate (either LoRa or FSK)
	Board             int           `json:"board"`                       // Concentrator board used for RX
	Antenna           int           `json:"antenna"`                     // Antenna number on which signal has been received
}

type Payload struct {
	RXInfo			RXInfo	`json:"rxInfo"`
	PyPayload		[]byte	`json:"phyPayLoad"`
}

type PyPayLoad struct {
	MHDR Mhdr	`json:"mhdr"`
	MacPayload	MacPayload `json:"macPayload"`

}


type Mhdr struct {
	MType 	string `json:"mType"`
	Major	string `json:"major"`
}


type MacPayload struct {
	FHDR	Fhdr `json:"fhdr"`
}

type Fhdr struct {
	DevAddr string `json:"devAddr"`
}


type MessageInfo struct {
	MAC lorawan.EUI64
	DevAddr	string
	Type	string
}

func createConnection() (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}


func run() error {
	client, err := createConnection()

	channel1 := make(chan MessageInfo)


	if err != nil {
		return err
	}

	for {
		if token := client.Subscribe("#", 0, func(client mqtt.Client, msg mqtt.Message) {
			go decodeMessage(msg, channel1)

			select {
				case t1 := <-channel1:
					fmt.Println("Message", t1)
			}

		}); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}

}

func main() {

	//run()
	var devEUI lorawan.EUI64
	copy(devEUI[:], []byte{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})

	pool := redis.NewRedisPool("redis://localhost:6379")
	_, err := redis.GetDeviceSession(pool, devEUI)

	if err != nil {
		fmt.Println(err)
	}

}

func handleMessage(message MessageInfo) {

}

func getKeys(devaddr string) {

}

func forwardMessage() {

}


func discardAndLogMessage() {

}

func decodeMessage(msg mqtt.Message, c1 chan MessageInfo)  (*MessageInfo, error) {
	var phy lorawan.PHYPayload
	var test Payload
	var t1 PyPayLoad

	fmt.Println(string(msg.Payload()))
	fmt.Println(msg.Topic())

	if strings.HasSuffix(msg.Topic(), "/stats") {
		return nil, nil
	}


	if err := json.Unmarshal(msg.Payload(), &test); err != nil {
		return nil, err
	}

	if err := phy.UnmarshalBinary(test.PyPayload); err != nil {
		return nil, err
	}

	data, _ := phy.MarshalText()

	if err := phy.UnmarshalText(data); err != nil {
		return nil, err
	}


	phyJSON, err := phy.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(phyJSON, &t1); err != nil {
		return nil, err
	}

	fmt.Println(test.RXInfo.MAC)
	fmt.Println(t1.MHDR.MType)
	fmt.Println(t1.MacPayload.FHDR.DevAddr)

	c1 <- MessageInfo{
		MAC:test.RXInfo.MAC,
		DevAddr:t1.MacPayload.FHDR.DevAddr,
		Type:t1.MHDR.MType,
	}


	return &MessageInfo{
		MAC:test.RXInfo.MAC,
		DevAddr:t1.MacPayload.FHDR.DevAddr,
		Type:t1.MHDR.MType,
	}, nil
}

