package db

type Descs struct {
	ID          int `json:"id" form:"id"`
	Id_Type     int `json:"id_type" form:"id_type"`
	Id_Desain   int `json:"id_desain" form:"id_desain"`
	Id_Tampilan int `json:"id_tampilan" form:"id_tampilan"`
	Id_Kinerja  int `json:"id_kinerja" form:"id_kinerja"`
	Id_Kamera   int `json:"id_kamera" form:"id_kamera"`
	Id_Audio    int `json:"id_audio" form:"id_audio"`
	Id_Baterai  int `json:"id_baterai" form:"id_baterai"`
	Id_Fitur    int `json:"id_fitur" form:"id_fitur"`
}
