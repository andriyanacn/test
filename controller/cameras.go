package controller

//import packages
import (
	"net/http"
	"phone_review/config"
	"phone_review/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data cameras
func GetCameras(c echo.Context) error {
	var cameras []db.Cameras

	if err := config.DB.Find(&cameras).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Cameras Smartphone",
		"cameras": cameras,
	})
}

//Fungsi get cameras by ID
func GetCamerasByID(c echo.Context) error {
	var cameras db.Cameras
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&cameras).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Cameras Smartphone ",
		"cameras": cameras,
	})
}

//fungsi create new cameras
func CreateCameras(e echo.Context) error {
	cameras := db.Cameras{}
	e.Bind(&cameras)

	if err := config.DB.Save(&cameras).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data cameras",
		"cameras": cameras,
	})
}

//Fungsi Update Tabel Cameras
func UpdateCamerasByID(e echo.Context) error {
	cameras := db.Cameras{}
	e.Bind(&cameras)

	if err := config.DB.Updates(&cameras).Where("id= ?", cameras.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Cameras",
		"cameras": cameras,
	})
}

//Fungsi hapus data cameras
func DeleteCamerasByID(e echo.Context) error {
	var cameras db.Cameras
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&cameras)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"cameras": cameras,
		"message": "Data Berhasil Dihapus",
	})
}
