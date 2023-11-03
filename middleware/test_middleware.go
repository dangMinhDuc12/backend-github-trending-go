package middleware

import (
	"example/backend-github-trending/model"
	"example/backend-github-trending/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func TestMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			req := req.ReqSignIn{}

			//----- Start Bind user request to req variable -----//
			if err := c.Bind(&req); err != nil {
				log.Error(err)

				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message: err.Error(),
					Data: nil,
				})
			}


			//----- End Bind user request to req variable -----//

			// if req.Email != "dangminhduca3@gmail.com" {
			// 	return c.JSON(http.StatusBadRequest, model.Response{
			// 		StatusCode: http.StatusBadRequest,
			// 		Message: "This user is not admin",
			// 		Data: nil,
			// 	})
			// }

			//request only bind once time, so we have to set request in context and get them in handler
			c.Set("requestBody", req)
			return next(c)
		}
	}
}