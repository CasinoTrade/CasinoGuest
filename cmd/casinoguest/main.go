package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "none"
)

const (
	httpPort = "8080"
	baseURL  = ":" + httpPort
)

func main() {
	printVersion := flag.Bool("version", false, "Get version")
	flag.Parse()
	if *printVersion {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Date: %s\n", date)
		fmt.Printf("Build Commit: %s\n", commit)
		return
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Number int }{Number: 42})
	})

	e.Logger.Fatal(e.Start(baseURL))
}
