package db

type Displays struct {
	ID           int    `json:"id" form:"id"`
	DimensiLayar string `json:"dimensilayar" form:"dimensilayar"`
	JenisLayar   string `json:"jenislayar" form:"jenislayar"`
	Pelindung    string `json:"pelindung" form:"pelindung"`
	Resolusi     string `json:"resolusi" form:"resolusi"`
}
