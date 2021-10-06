package middleware

import (
	"phone_review/config"
	"phone_review/db"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(adminname, password string, c echo.Context) (bool, error) {
	var admin db.Admins

	err := config.DB.Where("email = ? AND password = ?", adminname, password).First(&admin).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
