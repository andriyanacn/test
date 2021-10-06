package controller

//import packages
import (
	"net/http"
	"phone_review/config"
	"phone_review/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data desain
func GetDesains(c echo.Context) error {
	var desain []db.Desains

	if err := config.DB.Find(&desain).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Desain Smartphone",
		"desain":  desain,
	})
}

//Fungsi get desain by ID
func GetDesainsByID(c echo.Context) error {
	var desain db.Desains
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&desain).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan desain Smartphone ",
		"desain":  desain,
	})
}

//fungsi create new desain
func CreateDesains(e echo.Context) error {
	desain := db.Desains{}
	e.Bind(&desain)

	if err := config.DB.Save(&desain).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil menambahkan desain smartphone",
		"desain":  desain,
	})
}

//Fungsi Update Tabel desain
func UpdateDesainByID(e echo.Context) error {
	desain := db.Desains{}
	e.Bind(&desain)

	if err := config.DB.Updates(&desain).Where("id= ?", desain.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data tipe",
		"desain":  desain,
	})
}

//Fungsi hapus data desain
func DeleteDesainByID(e echo.Context) error {
	var desain db.Desains
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&desain)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"desain":  desain,
		"message": "Data Berhasil Dihapus",
	})
}
