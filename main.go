package main

import "github.com/labstack/echo/v4"

func main() {
	server := echo.New()
	server.GET("/", func(context echo.Context) error {
		return context.String(200, "Hello World")
	})
	server.Start("0.0.0.0:8080")
}
