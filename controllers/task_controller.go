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

func CreateTask(c echo.Context) error {
	validate := validator.New()

	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	if err := config.DB.Create(&task).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := validate.Struct(task); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Create task success")
}

func GetTasks(c echo.Context) error {
	var tasks []models.Task
	if err := config.DB.Find(&tasks).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Get task success")
}

func GetTaskByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, http.StatusNotFound, "Task not found")
		}
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Get task success")
}

func UpdateTask(c echo.Context) error {
	validate := validator.New()

	id, _ := strconv.Atoi(c.Param("id"))
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	if err := validate.Struct(task); err != nil {
		return helpers.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	existingTask := models.Task{}
	if err := config.DB.First(&existingTask, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, http.StatusNotFound, "Task not found")
		}
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Status = task.Status

	if err := config.DB.Save(&existingTask).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Update task success")
}

func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, http.StatusNotFound, "Task not found")
		}
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		return helpers.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, nil, "Delete task success")
}
