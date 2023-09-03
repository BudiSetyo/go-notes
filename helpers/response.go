package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorResponse adalah struktur yang digunakan untuk mengembalikan respons kesalahan JSON.
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// SuccessResponse adalah struktur yang digunakan untuk mengembalikan respons sukses JSON.
type SuccessResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendSuccessResponse mengirimkan respons sukses dengan status HTTP 200 OK.
func SendSuccessResponse(c echo.Context, data interface{}, message string) error {
	response := SuccessResponse{
		Error:   false,
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

// SendErrorResponse mengirimkan respons kesalahan dengan status HTTP yang diberikan.
func SendErrorResponse(c echo.Context, statusCode int, message string) error {
	response := ErrorResponse{
		Error:   true,
		Message: message,
	}
	return c.JSON(statusCode, response)
}
