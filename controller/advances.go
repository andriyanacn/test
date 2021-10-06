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
func GetAdvances(c echo.Context) error {
	var advances []db.Advances

	if err := config.DB.Find(&advances).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Berhasil Menampilkan Semua Kelebihan & Kekurangan Smartphone",
		"advances": advances,
	})
}

//Fungsi get admin by ID
func GetAdvancesByID(c echo.Context) error {
	var advances db.Advances
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&advances).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Berhasil Menampilkan Kelebihan & Kekurangan Smartphone ",
		"advances": advances,
	})
}

//fungsi create new admins
func CreateAdvances(e echo.Context) error {
	advances := db.Advances{}
	e.Bind(&advances)

	if err := config.DB.Save(&advances).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success menambahkan merks",
		"advances": advances,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateAdvanceByID(e echo.Context) error {
	advance := db.Advances{}
	e.Bind(&advance)

	if err := config.DB.Updates(&advance).Where("id= ?", advance.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Berhasil Mengubah Data advances",
		"advances": advance,
	})
}

//Fungsi hapus data merk
func DeleteAdvanceByID(e echo.Context) error {
	var advances db.Advances
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&advances)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"advances": advances,
		"message":  "Data Berhasil Dihapus",
	})
}
