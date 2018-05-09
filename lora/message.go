package lora

import (
	"time"
)


type MessageTest struct {
	TxInfo 	   *TxInfo1 `json:"txInfo"`
	RxInfo	   *RxInfo1 `json:"rxInfo"`
	PhyPayload string  `json:"phyPayload"`
}

type TxInfo1 struct {
	DataRate	*DateRate1	`json:"dataRate"`
	Frequency	int			`json:"frequency"`
	Adr			bool		`json:"adr"`
	CodeRate	string		`json:"codeRate"`
}

type DateRate1 struct {
	Modulation   string `json:"modulation"`
	Bandwidth    int    `json:"bandwidth"`
	SpreadFactor int 	`json:"spreadFactor"`
}

type Message struct {
	PhyPayload string  `json:"phyPayload"`
	RxInfo     *RxInfo `json:"rxInfo"`
}



type RxInfo1 struct {

	Mac 		string	`json:"mac"`
	Name 		string	`json:"name"`
	Latitude 	int		`json:"latitude"`
	Longitude 	int		`json:"longitude"`
	Altitude 	int		`json:"altitude"`
	Time 		string	`json:"time"`
	Rssi 		int		`json:"rssi"`
	LoRaSNR		int		`json:"loRaSNR"`
}

type RxInfo struct {
	Channel   int       `json:"channel"`
	CodeRate  string    `json:"codeRate"`
	CrcStatus int       `json:"crcStatus"`
	DataRate  *DataRate `json:"dataRate"`
	Frequency int       `json:"frequency"`
	LoRaSNR   int       `json:"loRaSNR"`
	Mac       string    `json:"mac"`
	RfChain   int       `json:"rfChain"`
	Rssi      int       `json:"rssi"`
	Size      int       `json:"size"`
	Time      string    `json:"time"`
	Timestamp int32     `json:"timestamp"`
}

type DataRate struct {
	Bandwidth    int    `json:"bandwidth"`
	Modulation   string `json:"modulation"`
	SpreadFactor int    `json:"spreadFactor"`
	BitRate      int    `json:"bitrate"`
}


func TestMessage(payload []byte, gwMac string, lat int) *MessageTest {


	dateRate := &DateRate1{
		Modulation:		"LORA",
		Bandwidth:		250,
		SpreadFactor:	5,
	}

	txInfo := &TxInfo1{
		Frequency:	868100000,
		DataRate:dateRate,
		Adr:false,
		CodeRate:"4/6",
	}


	rxInfo := &RxInfo1{
		Mac:		gwMac,
		Name:		"Hello",
		Latitude: 	lat,
		Longitude: 	4,
		Altitude:	10,
		Time:		"2018-01-20T11:24:37.295915988Z",
		Rssi:		-57,
		LoRaSNR: 	10,
	}

	message := &MessageTest{
		PhyPayload: string(payload),
		RxInfo:rxInfo,
		TxInfo:txInfo,
	}

	return message

}

func GenerateMessage(payload []byte, gwMac string) *Message {

	dataRate := &DataRate{
		Bandwidth:    250,
		Modulation:   "LORA",
		SpreadFactor: 5,
		BitRate:      0}

	rxInfo := &RxInfo{
		Channel:   0,
		CodeRate:  "4/6",
		CrcStatus: 1,
		DataRate:  dataRate,
		Frequency: 868100000,
		LoRaSNR:   7,
		Mac:       gwMac,
		RfChain:   1,
		Rssi:      -57,
		Size:      23,
		Time:      time.Now().Format(time.RFC3339),
		Timestamp: int32(time.Now().UnixNano() / 1000000000)}

	message := &Message{
		PhyPayload: string(payload),
		RxInfo:     rxInfo}

	return message
}



