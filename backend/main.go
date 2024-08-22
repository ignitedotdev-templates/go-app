package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		appVersion := os.Getenv("APP_VERSION")

		response := map[string]interface{}{
			"message":     "This is your APP_VERSION environment variable.",
			"app_version": appVersion,
		}

		return c.JSON(http.StatusOK, response)
	})

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
