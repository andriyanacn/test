package controller

//import packages
import (
	"net/http"
	"phone_review/config"
	"phone_review/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data admin
func GetBatteries(c echo.Context) error {
	var battery []db.Batteries

	if err := config.DB.Find(&battery).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Baterai Smartphone",
		"battery": battery,
	})
}

//Fungsi get admin by ID
func GetBatteriesByID(c echo.Context) error {
	var battery db.Batteries
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&battery).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Baterai Smartphone ",
		"battery": battery,
	})
}

//fungsi create new admins
func CreateBattery(e echo.Context) error {
	battery := db.Batteries{}
	e.Bind(&battery)

	if err := config.DB.Save(&battery).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data battery",
		"battery": battery,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateBatteryByID(e echo.Context) error {
	battery := db.Batteries{}
	e.Bind(&battery)

	if err := config.DB.Updates(&battery).Where("id= ?", battery.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Baterai",
		"battery": battery,
	})
}

//Fungsi hapus data merk
func DeleteBatteryByID(e echo.Context) error {
	var battery db.Batteries
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&battery)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"battery": battery,
		"message": "Data Berhasil Dihapus",
	})
}
