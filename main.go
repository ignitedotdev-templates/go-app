package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	e := echo.New()

	var users []User
	users = append(users, User{Name: "Precious", Age: 10})
	users = append(users, User{Name: "Ijeoma", Age: 20})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
