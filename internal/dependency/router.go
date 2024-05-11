package dependency

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rarrazaan/be-player-performance-app/internal/middleware"
)

func NewRouter(dep Dependency, ddep DirectDependency) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery()).
		Use(requestid.New()).
		Use(middleware.WithTimeout, middleware.GlobalErrorMiddleware())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	})

	return r
}
