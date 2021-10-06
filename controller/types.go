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
func GetTypes(c echo.Context) error {
	var types []db.Types

	if err := config.DB.Find(&types).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Tipe Smartphone",
		"types":   types,
	})
}

//Fungsi get admin by ID
func GetTypesByID(c echo.Context) error {
	var types db.Types
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&types).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Tipe Smartphone ",
		"types":   types,
	})
}

//fungsi create new admins
func CreateType(e echo.Context) error {
	tipe := db.Types{}
	e.Bind(&tipe)

	if err := config.DB.Save(&tipe).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil menambahkan tipe smartphone",
		"tipe":    tipe,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateTypeByID(e echo.Context) error {
	tipe := db.Types{}
	e.Bind(&tipe)

	if err := config.DB.Updates(&tipe).Where("id= ?", tipe.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data tipe",
		"tipe":    tipe,
	})
}

//Fungsi hapus data merk
func DeleteTypeByID(e echo.Context) error {
	var tipe db.Types
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&tipe)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"tipe":    tipe,
		"message": "Data Berhasil Dihapus",
	})
}
