package db

type Types struct {
	ID       int    `json:"id" form:"id"`
	Id_Merk  int    `json:"id_merk" form:"id_merk"`
	Tipe     string `json:"tipe" form:"tipe"`
	Th_Rilis int    `json:"th_rilis" form:"th_rilis"`
	Harga    int    `json:"harga" form:"harga"`
}
