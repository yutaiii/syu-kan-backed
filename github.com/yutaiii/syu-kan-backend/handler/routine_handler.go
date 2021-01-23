package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yutaiii/syu-kan-backend/usecase"
)

func GetRoutines() echo.HandlerFunc {
	return func(c echo.Context) error {
		routines, err := usecase.GetAllRoutines()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, routines)
	}
}
