package controller

import (
	"fmt"
	"net/http"
	httpEntity "yaya_short_url_local/src/entity/http"
	services "yaya_short_url_local/src/service"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	UrlService    services.UrlServiceInterface
	InsertService services.InsertInterface
}

func (service *UrlController) ShortenUrl(context *gin.Context) {
	var form httpEntity.UrlForm
	var res httpEntity.Response
	context.BindJSON(&form)

	if form.URL == "" {
		res.Success = false
		res.Message = "url is not present"
		context.JSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println(form)
	result, err, message, status := service.UrlService.ShortenUrl(form)
	fmt.Println(result)
	if err != nil || result == nil {
		res.Success = false
		res.Message = message
		context.JSON(status, res)
		return
	}
	res.Success = true
	res.Message = message
	res.Data = result
	context.JSON(status, res)
}

func (service *UrlController) GetShortCode(context *gin.Context) {

	var res httpEntity.Response

	urlShortcode := context.Param("shortcode")

	result, err, message, status := service.UrlService.GetShortCode(urlShortcode)
	if nil != err || result.ID == 0 {
		res.Success = false
		res.Message = message
		context.JSON(status, res)
		return
	}
	context.Redirect(http.StatusFound, result.Url)
	return

}

func (service *UrlController) GetStats(context *gin.Context) {

	var res httpEntity.Response

	urlShortcode := context.Param("shortcode")

	result, err, message, status := service.UrlService.GetStats(urlShortcode)
	if nil != err || result == nil {
		res.Success = false
		res.Message = message
		context.JSON(status, res)
		return
	}
	res.Success = true
	res.Message = message
	res.Data = result
	context.JSON(status, res)

}
