package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/usecase"
)

func CreateProgerss() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new([]*model.Progress)
		err := c.Bind(p)
		if err != nil {
			log.Printf("CreateProgerss, Bind err: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		err = usecase.CreateProgress(*p)
		if err != nil {
			log.Printf("CreateProgerss err: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "OK")
	}
}
