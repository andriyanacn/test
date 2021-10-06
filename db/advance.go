package db

type Advances struct {
	ID         int    `json:"id" form:"id"`
	Type_Id    int    `json:"type_id" form:"type_id"`
	Kelebihan  string `json:"kelebihan" form:"kelebihan"`
	Kekurangan string `json:"kekurangan" form:"kekurangan"`
}
