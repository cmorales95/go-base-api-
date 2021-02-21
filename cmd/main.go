package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/crisdavidmm95.dev/aug7/internal/app/database"
)

type User struct {
	att int
}

func main() {

	database.ConfigureDatabaseSQL()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world from Aug7 App!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
