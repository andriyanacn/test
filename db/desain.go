package db

type Desains struct {
	ID           int    `json:"id" form:"id"`
	BodyMaterial string `json:"bodymaterial" form:"bodymaterial"`
	Dimensi      string `json:"dimensi" form:"dimensi"`
	Ketebalan    string `json:"ketebalan" form:"ketebalan"`
	Berat        string `json:"berat" form:"berat"`
}
