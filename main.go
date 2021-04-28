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

	// routines
	// 使用していないAPIをコメントアウト
	// e.GET("/routines", handler.GetRoutines())
	e.POST("/routines/create", handler.CreateRoutines())
	e.PUT("/users/:userId/routines", handler.UpdateRoutines())
	e.DELETE("/users/:userId/routines", handler.DeleteRoutines())

	// progress
	e.POST("/progress", handler.CreateProgerss())
	e.GET("/users/:userId/progress/today", handler.GetProgressOfToday())

	// users
	e.POST("/user/create", handler.CreateUser())
	e.GET("/users/:userId/routines", handler.GetRoutinesByUserId())
	e.POST("/users/find", handler.FindUserByFirebaseUID())

	//exec server
	e.Start(":8000")
}
