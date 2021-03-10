package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

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

func GetRoutinesByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase := usecase.NewRoutineUsecase(context.Background())
		userId := c.Param("userId")
		var m model.RoutineForGetInput
		uintUserId, err := strconv.ParseUint(userId, 10, 64)
		if err != nil {
			log.Printf("RoutineAPI, strconv.ParseUint error: %+v", err)
			return c.JSON(http.StatusInternalServerError, "error")
		}

		m.UserID = uintUserId
		routines, err := usecase.GetByUserId(&m)
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

func UpdateRoutines() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase := usecase.NewRoutineUsecase(context.Background())
		var models []*model.Routine
		err := c.Bind(&models)
		if err != nil {
			log.Printf("RoutineAPI, UpdateRoutines, Bind error: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		// TODO 別メソッドに切り出す
		userIdStr := c.Param("userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			log.Printf("RoutineAPI, UpdateRoutines, strconv.ParseInt error: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		for i := 0; i < len(models); i++ {
			models[i].UserID = uint64(userId)
		}

		result, err := usecase.UpdateRoutines(models)
		if err != nil {
			log.Printf("RoutineAPI, UpdateRoutines error: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func DeleteRoutines() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase := usecase.NewRoutineUsecase(context.Background())
		models := new([]*model.Routine)
		err := c.Bind(models)
		if err != nil {
			log.Printf("RoutineAPI, DeleteRoutines,Bind error: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		err = usecase.DeleteRoutines(*models)
		if err != nil {
			log.Printf("RoutineAPI, DeleteRoutines error: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "OK")
	}
}
