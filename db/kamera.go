package db

type Cameras struct {
	ID          int    `json:"id" form:"id"`
	Jml_Blkg    int    `json:"jml_blkg" form:"jml_blkg"`
	Jml_Depan   int    `json:"jml_depan" form:"jml_depan"`
	MainCamera  string `json:"maincamera" form:"maincamera"`
	FrontCamera string `json:"selfiecamera" form:"selfiecamera"`
	LEDFlash    string `json:"ledflash" form:"ledflash"`
	AutoFocus   string `json:"autofocus" form:"autofocus"`
}
