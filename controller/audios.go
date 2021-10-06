package controller

//import packages
import (
	"net/http"
	"phone_review/config"
	"phone_review/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data audio
func GetAudios(c echo.Context) error {
	var audios []db.Audios

	if err := config.DB.Find(&audios).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Audio Smartphone",
		"audios":  audios,
	})
}

//Fungsi get audio by ID
func GetAudiosByID(c echo.Context) error {
	var audios db.Audios
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&audios).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Audio Smartphone ",
		"audios":  audios,
	})
}

//fungsi create new audio
func CreateAudio(e echo.Context) error {
	audio := db.Audios{}
	e.Bind(&audio)

	if err := config.DB.Save(&audio).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data audio",
		"audio":   audio,
	})
}

//Fungsi Update Tabel Audio
func UpdateAudioByID(e echo.Context) error {
	audio := db.Audios{}
	e.Bind(&audio)

	if err := config.DB.Updates(&audio).Where("id= ?", audio.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Audio",
		"audio":   audio,
	})
}

//Fungsi hapus data audio
func DeleteAudioByID(e echo.Context) error {
	var audio db.Audios
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&audio)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"audio":   audio,
		"message": "Data Berhasil Dihapus",
	})
}
