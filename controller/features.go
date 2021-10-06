package controller

//import packages
import (
	"net/http"
	"phone_review/config"
	"phone_review/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data feature
func GetFeatures(c echo.Context) error {
	var feature []db.Features

	if err := config.DB.Find(&feature).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Features Smartphone",
		"feature": feature,
	})
}

//Fungsi get feature by ID
func GetFeaturesByID(c echo.Context) error {
	var feature db.Features
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&feature).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Features Smartphone ",
		"feature": feature,
	})
}

//fungsi create new feature
func CreateFeature(e echo.Context) error {
	feature := db.Features{}
	e.Bind(&feature)

	if err := config.DB.Save(&feature).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data feature",
		"feature": feature,
	})
}

//Fungsi Update Tabel Feature
func UpdateFeatureByID(e echo.Context) error {
	feature := db.Features{}
	e.Bind(&feature)

	if err := config.DB.Updates(&feature).Where("id= ?", feature.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Feature",
		"feature": feature,
	})
}

//Fungsi hapus Features data
func DeleteFeatureByID(e echo.Context) error {
	var feature db.Features
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&feature)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"feature": feature,
		"message": "Data Berhasil Dihapus",
	})
}
