package routes

import (
	"github.com/labstack/echo"

	"github.com/droxey/gogogadget/resources"
)

// RegisterUser creates a user and persists the data.
func RegisterUser(e *echo.Echo) {
	userController := new(resources.UserController)

	e.POST("auth/register", userController.Register)
	e.POST("auth/login", userController.Login)

}
