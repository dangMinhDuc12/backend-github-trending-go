package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func   HandleSignin(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "Duc",
		"email": "duc@gmail.com",
	})
}

func HandleSignup(c echo.Context) error {
	type User struct {
		Email string `json:"emailMap"` //thay đổi tên trường Email mà hàm này sẽ trả ra thành emailMap
		FullName string
		Age int
	}

	user := User{
		Email: "duc@gmail.com",
		FullName: "DucDang",
		Age: 10,
	}

	return c.JSON(http.StatusOK, user)
}