package middleware

import (
	"example/backend-github-trending/model"
	"example/backend-github-trending/security"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var token string


			 for key, values := range c.Request().Header {
						if key == "Authorization" {
     					for _,value := range values {
								splitString := strings.Split(value, " ")
                token = splitString[1]
            	}
							break
						}
        }

		claim, err := security.ParseToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, model.Response{
				StatusCode: http.StatusUnauthorized,
				Message: "Token không hợp lệ",
				Data: nil,
			})
		}

			c.Set("userContext", claim)

			return next(c)
		}
	}
}