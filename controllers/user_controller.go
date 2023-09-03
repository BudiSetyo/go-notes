package controllers

import (
	"net/http"
	"strconv"

	"go-notes/config"
	"go-notes/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "User not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
