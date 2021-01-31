package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/yutaiii/syu-kan-backend/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// Allows requests from any `http://localhost:8080` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	//routing
	e.GET("/routines", handler.GetRoutines())
	e.POST("/routines/create", handler.CreateRoutines())
	e.PUT("/routines/update", handler.UpdateRoutines())
	e.DELETE("/routines/delete", handler.DeleteRoutines())
	e.POST("/progress", handler.CreateProgerss())

	//exec server
	e.Start(":8000")
}
