package route

import (
	"amitshekar-clean-arch/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api/products")
	// All Public APIs
	NewProductRouter(env, db, publicRouter)

	// protectedRouter := gin.Group("/v1")

	// Middleware to verify AccessToken
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Private APIs
	// NewProfileRouter(env, timeout, db, protectedRouter)
}
