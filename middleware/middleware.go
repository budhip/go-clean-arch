package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

const (
	devHeader = `Api-App-Key`
)

// CustomMiddleware type
type CustomMiddleware struct {
}

// ResponseError type struct
type ResponseError struct {
	Message string `json:"message"`
}

// CheckAuthHeader function for checking API KEY
func (cm *CustomMiddleware) CheckAuthHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authDevHeader := c.Request().Header.Get(devHeader)
		apiKey := os.Getenv("API_KEY")

		if authDevHeader == apiKey {
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, &ResponseError{Message: `You are not authorized`})
	}

}

// InitMiddleware initial middleware
func InitMiddleware() *CustomMiddleware {
	return &CustomMiddleware{}
}
