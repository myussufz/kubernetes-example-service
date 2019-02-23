package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World")
		return c.String(200, fmt.Sprintf("%s from pod %s", os.Getenv("SYSTEM_NAME"), os.Getenv("HOSTNAME")))
	})

	e.Logger.Fatal(e.Start(":5000"))
}
