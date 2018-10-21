package main

import (
	"net/http"
	"github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("*", func(c echo.Context) error {
        st := text()
        st_b := text_b()
		return c.String(http.StatusOK, "Hello, World!\n" + st +"\n" + st_b)
	})

	// Start server
    e.Logger.Fatal(e.Start(":8520"))
}