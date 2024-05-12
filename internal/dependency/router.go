package dependency

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rarrazaan/be-player-performance-app/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(dep Dependency, ddep DirectDependency) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery()).
		Use(requestid.New()).
		Use(middleware.WithTimeout, middleware.GlobalErrorMiddleware(), middleware.CORSMiddleware())

	r.GET("/oauth/google", ddep.AuthHandler.GoogleLogin).
		GET("/oauth/google-callback", ddep.AuthHandler.GoogleLoginCallback).
		GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)).
		GET("/login-success", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Login Success"})
		})

	api := r.Group("/api", middleware.Auth(dep.Config))
	api.GET("/identity", ddep.IdentityPerformanceHandler.Identity).
		GET("/calculate", ddep.IdentityPerformanceHandler.CalculatePerformance)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	})

	return r
}
