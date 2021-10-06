package db

type Features struct {
	ID          int    `json:"id" form:"id"`
	Bluetooth   string `json:"bluetooth" form:"bluetooth"`
	InfraRed    string `json:"infrared" form:"infrared"`
	WiFi        string `json:"wifi" form:"wifi"`
	FingerPrint string `json:"fingerprint" form:"fingerprint"`
	NFC         string `json:"nfc" form:"nfc"`
}
