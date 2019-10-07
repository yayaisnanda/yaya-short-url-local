package controller

import (
	srv "yaya_short_url_local/src/service"

	"github.com/gin-gonic/gin"
)

func LoadRouter(routers *gin.Engine) {
	router := &UrlRouterLoader{}
	router.UrlRouter(routers)
}

type UrlRouterLoader struct {
}

func (rLoader *UrlRouterLoader) UrlRouter(router *gin.Engine) {
	handler := &UrlController{
		UrlService: srv.UrlServiceHandler(),
	}
	rLoader.routerDefinition(router, handler)
}

func (rLoader *UrlRouterLoader) routerDefinition(router *gin.Engine, handler *UrlController) {
	//router.POST("/shorten", handler.ShortenUrl)
	//router.GET("/:shortcode", handler.GetShortCode)
	//router.GET("/:shortcode/stats", handler.GetStats)

	router.POST("/shorten", handler.ShortenUrl)
	router.GET("/:shortcode", handler.GetShortCode)
	router.GET("/:shortcode/stats", handler.GetStats)

}
