package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/usecase"
)

func GetRoutines() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase := usecase.NewRoutineUsecase(context.Background())
		routines, err := usecase.GetAllRoutines()
		if err != nil {
			log.Printf("RoutineAPI, GetAllRoutines error: %+v", err)
			return c.JSON(http.StatusInternalServerError, "error")
		}
		return c.JSON(http.StatusOK, routines)
	}
}

func CreateRoutines() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase := usecase.NewRoutineUsecase(context.Background())
		models := new([]*model.Routine)
		err := c.Bind(models)
		if err != nil {
			log.Printf("RoutineAPI, Bind error: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		result, err := usecase.CreateRoutines(*models)
		if err != nil {
			log.Printf("RoutineAPI, CreateRoutines error: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, result)
	}
}
