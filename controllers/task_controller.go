package controllers

import (
	"net/http"
	"strconv"

	"go-notes/config"
	"go-notes/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := config.DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, task)
}

func GetTasks(c echo.Context) error {
	var tasks []models.Task
	if err := config.DB.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Task not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	existingTask := models.Task{}
	if err := config.DB.First(&existingTask, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Task not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Status = task.Status

	if err := config.DB.Save(&existingTask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, existingTask)
}

func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Task not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
