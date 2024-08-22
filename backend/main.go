package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		envVars := make(map[string]string)
		for _, env := range os.Environ() {
			pair := splitAtFirstEqual(env)
			envVars[pair[0]] = pair[1]
		}

		response := map[string]interface{}{
			"message": "These are your current environment variables. Update an env and refresh.",
			"envs":    envVars,
		}

		return c.JSON(http.StatusOK, response)
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("port is %v", port)
	e.Logger.Fatal(e.Start(port))
}

// splitAtFirstEqual splits a string at the first occurrence of '='
// and returns a slice with the key and value as separate elements.
func splitAtFirstEqual(s string) []string {
	for i := range s {
		if s[i] == '=' {
			return []string{s[:i], s[i+1:]}
		}
	}
	return []string{s, ""}
}
