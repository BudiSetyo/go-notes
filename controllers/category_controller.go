package controllers

import (
	"net/http"
	"strconv"

	"go-notes/config"
	"go-notes/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, category)
}

func GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, categories)
}

func GetCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Category not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	existingCategory := models.Category{}
	if err := config.DB.First(&existingCategory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Category not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	existingCategory.Name = category.Name

	if err := config.DB.Save(&existingCategory).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, existingCategory)
}

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Category not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
