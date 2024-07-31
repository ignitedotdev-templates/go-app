package main

import (
	"fmt"
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
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("port is %v", port)
	e.Logger.Fatal(e.Start(port))
}

//nixpacks build --name mn .  --install-cmd  "go mod download"  --build-cmd "go build -o out backend/main.go"
