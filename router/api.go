package router

import (
	"example/backend-github-trending/handler"

	"github.com/labstack/echo/v4"
)


type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
}

			//user router
func (api *API) SetupRouter() {
	api.Echo.GET("/user/sign-in", api.UserHandler.HandleSignin)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignup)
	api.Echo.GET("/user", api.UserHandler.HandleGetListUser)
}