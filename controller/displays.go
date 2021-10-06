package controller

//import packages
import (
	"net/http"
	"phone_review/config"
	"phone_review/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data displays
func GetDisplays(c echo.Context) error {
	var displays []db.Displays

	if err := config.DB.Find(&displays).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Berhasil Menampilkan Semua Data Displays Smartphone",
		"displays": displays,
	})
}

//Fungsi get displaysby ID
func GetDisplaysByID(c echo.Context) error {
	var displays db.Displays
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&displays).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Berhasil Menampilkan Displays Smartphone ",
		"displays": displays,
	})
}

//fungsi create new displays
func Createdisplays(e echo.Context) error {
	displays := db.Displays{}
	e.Bind(&displays)

	if err := config.DB.Save(&displays).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success menambahkan data feature",
		"displays": displays,
	})
}

//Fungsi Update Tabel Displays
func UpdateDisplaysByID(e echo.Context) error {
	displays := db.Features{}
	e.Bind(&displays)

	if err := config.DB.Updates(&displays).Where("id= ?", displays.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Berhasil Mengubah Data Displays",
		"displays": displays,
	})
}

//Fungsi hapus data displays
func DeleteDisplayByID(e echo.Context) error {
	var displays db.Displays
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&displays)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"displays": displays,
		"message":  "Data Berhasil Dihapus",
	})
}
