package db

type Audios struct {
	ID             int    `json:"id" form:"id"`
	Jml_Speaker    int    `json:"jml_speaker" form:"jml_speaker"`
	Jml_Microphone int    `json:"jml_microphone" form:"jml_microphone"`
	AudioJack      string `json:"audiojack" form:"audiojack"`
	Radio          string `json:"radio" form:"radio"`
}
