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

func GetProgressOfToday() echo.HandlerFunc {
	return func(c echo.Context) error {
		var m model.RoutineForGetInput

		userIdStr := c.Param("userId")
		userId, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			log.Printf("GetProgressOfToday, strconv.ParseInt err: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}
		m.UserID = userId

		usecase := usecase.NewProgressUsecase(context.Background())
		progress, err := usecase.GetProgressOfToday(&m)
		if err != nil {
			log.Printf("GetProgressOfToday, GetProgressOfToday err: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, progress)
	}
}

func CreateProgerss() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new([]*model.Progress)
		err := c.Bind(p)
		if err != nil {
			log.Printf("CreateProgerss, Bind err: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		usecase := usecase.NewProgressUsecase(context.Background())
		err = usecase.CreateProgress(*p)
		if err != nil {
			log.Printf("CreateProgerss err: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "OK")
	}
}
