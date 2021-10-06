package db

type Performs struct {
	ID         int    `json:"id" form:"id"`
	RAM        string `json:"ram" form:"ram"`
	ROM        string `json:"rom" form:"rom"`
	Chipset    string `json:"chipset" form:"chipset"`
	OS         string `json:"os" form:"os"`
	SlotMemory string `json:"slotmemory" form:"slotmemory"`
	SimCard    string `json:"sim_card" form:"sim_card"`
}
