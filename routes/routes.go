package routes

import (
	"go-notes/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/users", controllers.CreateUser)
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUserByID)

	e.POST("/tasks", controllers.CreateTask)
	e.GET("/tasks", controllers.GetTasks)
	e.GET("/tasks/:id", controllers.GetTaskByID)
	e.PUT("/tasks/:id", controllers.UpdateTask)
	e.DELETE("/tasks/:id", controllers.DeleteTask)

	e.POST("/categories", controllers.CreateCategory)
	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategoryByID)
	e.PUT("/categories/:id", controllers.UpdateCategory)
	e.DELETE("/categories/:id", controllers.DeleteCategory)
}
