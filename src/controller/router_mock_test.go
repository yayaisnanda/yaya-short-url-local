package controller

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func LoadRouterTestMock() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	resp := httptest.NewRecorder()
	context, routers := gin.CreateTestContext(resp)

	routerLoader := &UrlRouterLoader{}
	routerLoader.UrlRouterTestMock(routers)

	return context, routers, resp
}

func (rLoader *UrlRouterLoader) UrlRouterTestMock(router *gin.Engine) {
	handler := &UrlController{
		UrlService: &UrlServiceMock{},
	}
	rLoader.routerDefinition(router, handler)
}
