package config

import (
	"yaya_short_url_local/src/util/middleware"

	"github.com/gin-gonic/gin"
)

// default server router configuration
func SetupRouter() *gin.Engine {
	defaultMiddleware := middleware.DefaultMiddleware{}
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}
