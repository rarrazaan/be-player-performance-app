package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	errorapp "github.com/rarrazaan/be-player-performance-app/internal/utils/errors"
)

var mapErrorCode = map[errorapp.CustomErrorCode]int{
	errorapp.BadRequest:     http.StatusBadRequest,
	errorapp.Forbidden:      http.StatusForbidden,
	errorapp.Unauthorized:   http.StatusUnauthorized,
	errorapp.InternalServer: http.StatusInternalServerError,
	errorapp.NotFound:       http.StatusNotFound,
}

func GlobalErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *errorapp.CustomError:
				c.AbortWithStatusJSON(mapErrorCode[e.Code], e.CreateHTTPErrorMessage())
			case validator.ValidationErrors:
				errs := errorapp.ValidationErrResponse(e)
				c.AbortWithStatusJSON(http.StatusBadRequest, dto.JSONResponse{
					Message: errs,
				})
			default:
				if errors.Is(err, context.DeadlineExceeded) {
					c.AbortWithStatus(http.StatusRequestTimeout)
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, dto.JSONResponse{
						Message: "internal server error",
					})
				}
			}

			c.Abort()
		}
	}
}
