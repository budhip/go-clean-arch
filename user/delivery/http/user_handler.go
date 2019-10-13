package http

import (
	"net/http"
	"testgo/models"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/budhip/go-postgre-clean-arch/user"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type UserHandler struct {
	UUsecase user.Usecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewUserHandler(e *echo.Echo, us user.Usecase) {
	handler := &UserHandler{
		UUsecase: us,
	}
	e.GET("/users", handler.FetchUsers)
}

// FetchUsers will fetch all the user
func (handler *UserHandler) FetchUsers(c echo.Context) error {
	listUser, err := handler.UUsecase.FetchUsers()

	var result map[string]interface{}

	if err != nil {
		result = map[string]interface{}{
			"message": err.Error(),
			"data":    listUser,
		}
		return c.JSON(GetStatusCode(err), result)
	}

	result = map[string]interface{}{
		"message": "success",
		"data":    listUser,
	}

	return c.JSON(http.StatusOK, result)
}

// GetStatusCode represent the statusCode
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
