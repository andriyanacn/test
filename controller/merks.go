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
func GetMerks(c echo.Context) error {
	var merks []db.Merks

	if err := config.DB.Find(&merks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Merk Smartphone",
		"merks":   merks,
	})
}

//Fungsi get admin by ID
func GetMerksByID(c echo.Context) error {
	var merks db.Merks
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&merks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Merk Smartphone ",
		"merks":   merks,
	})
}

//fungsi create new admins
func CreateMerk(e echo.Context) error {
	merk := db.Merks{}
	e.Bind(&merk)

	if err := config.DB.Save(&merk).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan merks",
		"merk":    merk,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateMerkByID(e echo.Context) error {
	merk := db.Merks{}
	e.Bind(&merk)

	if err := config.DB.Updates(&merk).Where("id= ?", merk.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data merks",
		"merk":    merk,
	})
}

//Fungsi hapus data merk
func DeleteMerkByID(e echo.Context) error {
	var merk db.Merks
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&merk)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"merk":    merk,
		"message": "Data Berhasil Dihapus",
	})
}
