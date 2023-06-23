package main

import (
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// disable log to screen
	log.SetOutput(io.Discard)

	log.Println("Starting server on port 8080")
	server := echo.New()
	server.GET("/", func(context echo.Context) error {
		return context.String(200, "Hello World")
	})
	server.Start("0.0.0.0:8080")
}
