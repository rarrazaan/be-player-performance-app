package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"github.com/rarrazaan/be-player-performance-app/internal/constant"
)

func SetCookieAfterLogin(c *gin.Context, config config.Config, accessToken string) {
	accessTokenCookieExp := int(config.Jwt.AccessTokenExpiration) * 60

	if config.ServiceHost == "localhost" {
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie(constant.AccessTokenCookieName, accessToken, accessTokenCookieExp, "/", config.ServiceHost, false, true)
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(constant.AccessTokenCookieName, accessToken, accessTokenCookieExp, "/", config.ServiceHost, true, true)
}
