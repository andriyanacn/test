package controller

//import packages
import (
	"net/http"
	"phone_review/config"
	"phone_review/db"
	"phone_review/middleware"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data admin
func GetAdminsController(c echo.Context) error {
	var admins []db.Admins

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all admins",
		"admins":  admins,
	})
}

//Fungsi get admin by ID
func GetAdminsByID(c echo.Context) error {
	var admins db.Admins
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get admins ",
		"admins":  admins,
	})
}

//fungsi create new admins
func CreateAdminController(e echo.Context) error {
	admin := db.Admins{}
	e.Bind(&admin)

	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan admins",
		"admin":   admin,
	})
}

func UpdateAdminByID(e echo.Context) error {
	admin := db.Admins{}
	e.Bind(&admin)

	if err := config.DB.Updates(&admin).Where("id= ?", admin.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data admins",
		"admin":   admin,
	})
}

//Fungsi hapus data admin
func DeleteAdminByID(e echo.Context) error {
	var admin db.Admins
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&admin)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data Berhasil Dihapus",
	})
}

func LoginAdminController(e echo.Context) error {
	admin := db.Admins{}
	e.Bind(&admin)

	err := config.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}
	token, err := middleware.CreateToken(admin.ID, admin.Name)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	//adminRes := db.AdminRes{admin.ID, admin.Name, admin.Email, token}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"token":   token,
		//"user":    adminRes,
	})
}
