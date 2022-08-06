package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("middleware-1")
			return next(c)
		}
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("middleware-2")
			return next(c)
		}
	})

	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("pre middleware-1")
			return next(c)
		}
	})
	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("pre middleware-2")
			return next(c)
		}
	})

	// Routes
	e.GET("/", hello, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("middleware on router level")
			return next(c)
		}
	})

	// Start server
	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":1323")

}

func hello(c echo.Context) error {
	println("processing request")
	time.Sleep(5 * time.Second)
	return c.String(http.StatusOK, "Hello, World!")
}
