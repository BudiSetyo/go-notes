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

func CreateCategory(c echo.Context) error {
	validate := validator.New()

	category := new(models.Category)

	if err := validate.Struct(category); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(category); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Create category success")
}

func GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, categories, "Get categories success")
}

func GetCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, http.StatusNotFound, "Category not found")
		}
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, category, "Get category success")
}

func UpdateCategory(c echo.Context) error {
	validate := validator.New()

	id, _ := strconv.Atoi(c.Param("id"))
	category := new(models.Category)

	if err := validate.Struct(category); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(category); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	existingCategory := models.Category{}
	if err := config.DB.First(&existingCategory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, http.StatusNotFound, "Category not found")
		}
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	existingCategory.Name = category.Name

	if err := config.DB.Save(&existingCategory).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Update category success")
}

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, http.StatusNotFound, "Category not found")
		}
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Delete category success")
}
