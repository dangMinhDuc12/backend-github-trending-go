package main

import (
	"example/backend-github-trending/db"
	"example/backend-github-trending/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		UserName: "postgres",
		Password: "postgres",
		Host: "localhost",
		Port: 5432,
		DbName: "golang",
	}

	sql.Connect()

	defer sql.Close() //defer: pending hàm này cho đến khi các function khác của hàm main chạy xong


    e := echo.New()
    e.GET("/", handler.Welcome)

		//user router
		e.GET("/user/sign-in", handler.HandleSignin)
		e.GET("user/sign-up", handler.HandleSignup)

    e.Logger.Fatal(e.Start(":3000"))
}




