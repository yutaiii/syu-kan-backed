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
			return c.JSON(http.StatusInternalServerError, "error")
		}
		return c.JSON(http.StatusOK, routines)
	}
}
