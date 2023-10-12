package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			 for key, values := range c.Request().Header {
            fmt.Println("key", key)
            for _,value := range values {
                fmt.Println("value", value)
            }
        }

			return next(c)
		}
	}
}