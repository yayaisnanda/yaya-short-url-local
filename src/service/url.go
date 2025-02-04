package services

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
	dbEntity "yaya_short_url_local/src/entity/db"

	"github.com/jinzhu/copier"
	"github.com/lucasjones/reggen"

	httpEntity "yaya_short_url_local/src/entity/http"
	repository "yaya_short_url_local/src/repository/db"
)

type UrlService struct {
	urlRepository repository.UrlRepositoryInterface
}

func UrlServiceHandler() *UrlService {
	return &UrlService{
		urlRepository: repository.UrlRepositoryHandler(),
	}
}

type UrlServiceInterface interface {
	GetShortCode(shortcode string) (*httpEntity.Url, error, string, int)
	GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int)
	ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int)
}

type InsertInterface interface {
	insertUrlToDB(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int)
	GetUrlFromDB() ([]*dbEntity.Url, error)
}

func GenerateShortUrl() string {
	g, err := reggen.NewGenerator("^[0-9a-zA-Z_]{6}$")
	if err != nil {
		panic(err)
	}
	return g.Generate(0)
}

func CheckInsertShortUrl(shortcode string) bool {
	r := regexp.MustCompile(`^[0-9a-zA-Z_]{4,}$`)
	matches := r.FindAllString(shortcode, 1)
	if len(matches) > 0 {
		return true
	}
	return false
}

func CheckFormToDB(urlList []*dbEntity.Url, shorcode string) bool {
	for i, _ := range urlList {
		if urlList[i].UrlShort == shorcode {
			return true
		}
	}
	return false
}

func (service *UrlService) insertUrlToDB(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int) {
	var url dbEntity.Url
	now := time.Now()
	url.CreatedAt = now
	url.Url = form.URL
	url.UrlShort = form.Shortcode

	err := service.urlRepository.InsertUrl(&url)
	if err != nil {
		message := "can't get url list from database"
		return nil, err, message, http.StatusBadGateway
	}

	var result httpEntity.Url
	result.ID = url.ID
	result.CreatedAt = url.CreatedAt
	result.Url = url.Url
	result.UrlShort = url.UrlShort
	message := "created"
	return &result, nil, message, http.StatusCreated
}

func (service *UrlService) GetUrlFromDB() ([]*dbEntity.Url, error) {
	fmt.Print("get url from db")
	urlList, err := service.urlRepository.GetUrlList()
	if err != nil {
		return nil, err
	}
	return urlList, err
}

func (service *UrlService) ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int) {
	urlList, err := service.GetUrlFromDB()
	if err != nil {
		message := "can't get url list from database"
		return nil, err, message, http.StatusBadGateway
	}
	if len(urlList) == 0 && form.Shortcode == "" {
		form.Shortcode = GenerateShortUrl()
		result, error, message, status := service.insertUrlToDB(form)
		return result, error, message, status

	} else if len(urlList) == 0 && form.Shortcode != "" {
		if CheckInsertShortUrl(form.Shortcode) {
			result, error, message, status := service.insertUrlToDB(form)
			return result, error, message, status
		} else {
			//not pass regex
			message := "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{4,}$."
			return nil, nil, message, http.StatusUnprocessableEntity
		}

	} else if len(urlList) > 0 && form.Shortcode == "" {
		for {
			form.Shortcode = GenerateShortUrl()
			if !CheckFormToDB(urlList, form.Shortcode) {
				result, error, message, status := service.insertUrlToDB(form)
				return result, error, message, status
			}
		}
	} else if len(urlList) > 0 && form.Shortcode != "" {
		if CheckInsertShortUrl(form.Shortcode) && !CheckFormToDB(urlList, form.Shortcode) {
			result, error, message, status := service.insertUrlToDB(form)
			return result, error, message, status
		} else if CheckInsertShortUrl(form.Shortcode) && CheckFormToDB(urlList, form.Shortcode) {
			message := "The the desired shortcode is already in use. Shortcodes are case-sensitive."
			return nil, nil, message, http.StatusConflict
		} else {
			message := "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{4,}$."
			return nil, nil, message, http.StatusUnprocessableEntity
		}
	}
	return nil, nil, "", 0
}

func (service *UrlService) GetShortCode(shortcode string) (*httpEntity.Url, error, string, int) {

	url, err := service.urlRepository.UpdateUrl(shortcode)
	if err != nil {
		message := "The shortcode cannot be found in the system"
		return nil, err, message, http.StatusNotFound
	}
	result := httpEntity.Url{}
	copier.Copy(&result, url)
	message := "success get shortcode"
	return &result, err, message, http.StatusOK
}

func (service *UrlService) GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int) {
	var result httpEntity.ShortcodeStats
	url, err := service.urlRepository.GetStats(shortcode)
	if err != nil {
		message := "The shortcode cannot be found in the system"
		return &result, err, message, http.StatusNotFound
	}
	result.LastSeenDate = url.LastSeenDate
	result.RedirectCount = url.RedirectCount
	result.StartDate = url.CreatedAt
	message := "success get status"
	return &result, err, message, http.StatusOK
}
