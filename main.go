package main

import (
	"example/backend-github-trending/db"
	"example/backend-github-trending/handler"
	repoimpl "example/backend-github-trending/repository/repo_impl"
	"example/backend-github-trending/router"

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
		userHandler := handler.UserHandler{
			UserRepo: repoimpl.NewUserRepo(sql),
		}

		api := &router.API{
			Echo: e,
			UserHandler: userHandler,
		}

		api.SetupRouter()

    e.Logger.Fatal(e.Start(":3000"))
}




