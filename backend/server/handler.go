package server

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	Handler interface {
		Routes(c *gin.Engine)
	}

	HandlerFuncErr func(*gin.Context) error
)

func HandleResponse[T any](c *gin.Context, successStatus int, body T, err error) error {
	if err != nil {
		return err
	}

	if k := reflect.TypeOf(body).Kind(); (k == reflect.Array || k == reflect.Slice) &&
		reflect.ValueOf(body).Len() == 0 {
		return NotFound
	}

	c.JSON(successStatus, body)

	return nil
}

func HandleErr(fn HandlerFuncErr) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := fn(ctx)
		if err == nil {
			ctx.Next()
			return
		}

		ctx.JSON(MapErr(err))
	}
}

func MapErr(err error) (int, HttpErr) {
	switch err := err.(type) {
	case *pgconn.PgError:
		if err.Code == "23505" {
			return http.StatusBadRequest, HttpErr{
				Description: "cannot add more entries because actual is duplicated",
			}
		}
	case APIResponseErr:
		return err.Status, err.HttpErr()
	}
	return http.StatusInternalServerError, HttpErr{
		Description: err.Error(),
	}
}
