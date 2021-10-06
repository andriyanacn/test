package db

type Batteries struct {
	ID           int    `json:"id" form:"id"`
	Daya         string `json:"daya" form:"daya"`
	JenisSocket  string `json:"jenissocket" form:"jenissocket"`
	FastCharging string `json:"fastcharging" form:"fastcharging"`
}
