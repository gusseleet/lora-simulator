package lora

import (
	"encoding/hex"
	"lora_simulator/logging"
)


func GenerateAddr (config *LoraConfig) ([16]byte, [16]byte, [4]byte) {
	var devAddr ([4]byte)
	da, _ := hex.DecodeString(config.DevHexAddr)
	copy(devAddr[:], da[:])

	var nwkSKey ([16]byte)
	nk, _ := hex.DecodeString(config.NwsHexkey)
	copy(nwkSKey[:], nk[:])
	var appSKey ([16]byte)
	ak, _ := hex.DecodeString(config.AppHexKey)
	copy(appSKey[:], ak[:])


	logging.Log().Debug("Hex values:", "\n devAddr:" , da,"\n nwkSKey:" , nk ,"\n appSKey:", ak)
	return appSKey, nwkSKey, devAddr

}