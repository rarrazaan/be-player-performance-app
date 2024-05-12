package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"github.com/rarrazaan/be-player-performance-app/internal/constant"
	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	"github.com/rarrazaan/be-player-performance-app/internal/utils"
	"github.com/rarrazaan/be-player-performance-app/internal/utils/errors"
)

func Auth(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("ENV_MODE") == "testing" {
			c.Next()
			return
		}

		accessTokenStr, err := c.Cookie(constant.AccessTokenCookieName)
		if err != nil {
			e := errors.ErrAccessTokenExpired
			c.AbortWithStatusJSON(mapErrorCode[e.Code], e.CreateHTTPErrorMessage())
			return
		}

		token, err := utils.ValidateAccessToken(accessTokenStr, config)
		if err != nil {
			if e, ok := err.(*errors.CustomError); ok {
				c.AbortWithStatusJSON(mapErrorCode[e.Code], e.CreateHTTPErrorMessage())
				return
			}

			e := errors.ErrInvalidToken
			c.AbortWithStatusJSON(mapErrorCode[e.Code], e.CreateHTTPErrorMessage())
			return
		}

		claims, ok := token.Claims.(*utils.AccessJWTClaim)
		if !ok || !token.Valid {
			if err := token.Claims.Valid(); err != nil {
				if e, ok := err.(*errors.CustomError); ok {
					c.AbortWithStatusJSON(mapErrorCode[e.Code], e.CreateHTTPErrorMessage())
					return
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.JSONResponse{
					Message: "internal server error",
				})
				return
			}
			return
		}

		user := &utils.ContextData{
			UserID:    claims.UserID,
			UserEmail: claims.UserEmail,
			UserName:  claims.UserName,
		}

		c.Set(string(utils.RequesterCtxKey), user)

		c.Next()
	}
}
