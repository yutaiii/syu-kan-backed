package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/usecase"
)

func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(*model.User)
		err := c.Bind(u)
		if err != nil {
			log.Printf("CreateUser, Bind err: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		usecase := usecase.NewUserUsecase(context.Background())
		err = usecase.CreateUser(*u)
		if err != nil {
			log.Printf("CreateUser err: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "OK")
	}
}

func FindUserByFirebaseUID() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(model.User)
		err := c.Bind(u)
		if err != nil {
			log.Printf("FindUserByFirebaseUID, Bind err: %+v", err)
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		usecase := usecase.NewUserUsecase(context.Background())
		res, err := usecase.FindUserByFirebaseUID(u)
		if err != nil {
			log.Printf("FindUserByFirebaseUID err: %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, res)
	}
}
