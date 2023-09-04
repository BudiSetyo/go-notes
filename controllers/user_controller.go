package controllers

import (
	"net/http"
	"strconv"

	"go-notes/config"
	"go-notes/helpers"
	"go-notes/models"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	validate := validator.New()

	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	if err := validate.Struct(user); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Create User Success")
}

func GetUsers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helpers.SendSuccessResponse(c, users, "Get User Success")
}

func GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, http.StatusNotFound, "User not found")
		}
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, user, "Get User Success")
}
