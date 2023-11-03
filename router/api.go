package router

import (
	"example/backend-github-trending/handler"
	localMiddleware "example/backend-github-trending/middleware"

	"github.com/labstack/echo/v4"
)


type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
}

			//user router
func (api *API) SetupRouter() {

	//authen router
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignin, localMiddleware.TestMiddleware())
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignup)



	//user route
	api.Echo.GET("/user", api.UserHandler.HandleGetListUser, localMiddleware.CheckToken())
	api.Echo.GET("/user/profile", api.UserHandler.HandleGetProfileUser, localMiddleware.CheckToken())
	api.Echo.PUT("/user/profile", api.UserHandler.HandleUpdateUser, localMiddleware.CheckToken())
}