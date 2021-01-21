package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/yutaiii/syu-kan-backend/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routing
	e.GET("/", handler.HelloWorld())

	//exec server
	e.Start(":8000")
}
