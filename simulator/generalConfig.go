package simulator


type GeneralConfig struct {
	NwkSKey [16]byte
	AppSKey [16]byte
	DevAddr [4]byte
	FPort uint8
}

type GeneralConfigHex struct {
	Hex_nwkSkey string
	Hex_appSkey string
	Hex_devAddr string
}

