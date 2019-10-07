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
	// ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int)
	// GetShortCode(shortcode string) (*httpEntity.Url, error, string, int)
	// GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int)

	ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int)
	GetShortCode(shortcode string) (*httpEntity.Url, error, string, int)
	GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int)
}

// func (service *UrlService) ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int) {
// 	var url dbEntity.Url

// 	url.Url = form.URL
// 	url.UrlShort = form.Shortcode
// 	_, err, message, status := service.urlRepository.ShortenUrl(&url)
// 	result := httpEntity.Url{}
// 	copier.Copy(&result, url)
// 	return &result, err, message, status
// }

// func (service *UrlService) GetShortCode(shortcode string) (*httpEntity.Url, error, string, int) {
// 	//var url *dbEntity.Url
// 	url, err, message, status := service.urlRepository.GetShortCode(shortcode)
// 	result := httpEntity.Url{}
// 	copier.Copy(&result, url)
// 	return &result, err, message, status
// }

// func (service *UrlService) GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int) {
// 	var result httpEntity.ShortcodeStats
// 	url, err, message, status := service.urlRepository.GetStats(shortcode)
// 	if url != nil {
// 		result.LastSeenDate = url.LastSeenDate
// 		result.RedirectCount = url.RedirectCount
// 		result.StartDate = url.CreatedAt
// 		return &result, err, message, status
// 	}
// 	return nil, err, message, status
// }

func (service *UrlService) ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int) {
	urlList, err, message, status := service.urlRepository.GetUrlList()
	if err != nil {
		return nil, err, message, status
	}
	var url dbEntity.Url

	var result httpEntity.Url

	now := time.Now()

	if len(urlList) == 0 {
		if form.Shortcode == "" {
			g, err := reggen.NewGenerator("^[0-9a-zA-Z_]{6}$")
			if err != nil {
				panic(err)
			}
			generate := g.Generate(0)
			form.Shortcode = generate

			url.CreatedAt = now
			url.Url = form.URL
			url.UrlShort = form.Shortcode

			err, message, status := service.urlRepository.InsertUrl(&url)
			if err != nil {
				return nil, err, message, status
			}
			result.ID = url.ID
			result.CreatedAt = url.CreatedAt
			result.Url = url.Url
			result.UrlShort = url.UrlShort
			return &result, nil, message, http.StatusCreated
		} else {
			r := regexp.MustCompile(`^[0-9a-zA-Z_]{4,}$`)
			matches := r.FindAllString(form.Shortcode, 1)

			checkRegex := false
			if len(matches) > 0 {
				checkRegex = true
			}

			if checkRegex {
				//fmt.Println("if")
				form.Shortcode = matches[0]
				url.CreatedAt = now
				url.Url = form.URL
				url.UrlShort = form.Shortcode
				err, message, status := service.urlRepository.InsertUrl(&url)
				if err != nil {
					return nil, err, message, status
				}
				result.ID = url.ID
				result.CreatedAt = url.CreatedAt
				result.Url = url.Url
				result.UrlShort = url.UrlShort
				return &result, nil, message, http.StatusCreated
			} else {
				//not pass regex
				//fmt.Println("else")
				message := "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{4,}$."
				return nil, nil, message, http.StatusUnprocessableEntity
			}
		}
	} else {
		if form.Shortcode == "" {
			for {
				g, err := reggen.NewGenerator("^[0-9a-zA-Z_]{6}$")
				if err != nil {
					panic(err)
				}
				generate := g.Generate(0)

				check := false
				for i, _ := range urlList {
					if urlList[i].UrlShort == generate {
						check = true
					}
				}
				if !check {
					// inssert shorten code to database
					fmt.Println("else-if-if")
					form.Shortcode = generate

					url.CreatedAt = now
					url.Url = form.URL
					url.UrlShort = form.Shortcode

					err, message, status := service.urlRepository.InsertUrl(&url)
					if err != nil {
						return nil, err, message, status
					}
					result.ID = url.ID
					result.CreatedAt = url.CreatedAt
					result.Url = url.Url
					result.UrlShort = url.UrlShort
					return &result, nil, message, http.StatusCreated
				}
			}
		} else {
			r := regexp.MustCompile(`^[0-9a-zA-Z_]{4,}$`)
			matches := r.FindAllString(form.Shortcode, 1)
			checkRegex := false
			if len(matches) > 0 {
				checkRegex = true
			}
			if checkRegex {
				check := false
				for i, _ := range urlList {
					if urlList[i].UrlShort == form.Shortcode {
						check = true
					}
				}
				if !check {
					// fmt.Println("insert to db")
					// inssert shorten code to database
					form.Shortcode = matches[0]
					url.CreatedAt = now
					url.Url = form.URL
					url.UrlShort = form.Shortcode
					err, message, status := service.urlRepository.InsertUrl(&url)
					if err != nil {
						return nil, err, message, status
					}
					result.ID = url.ID
					result.CreatedAt = url.CreatedAt
					result.Url = url.Url
					result.UrlShort = url.UrlShort
					return &result, nil, message, http.StatusCreated
				} else {
					// shorcode exist in database
					message := "The the desired shortcode is already in use. Shortcodes are case-sensitive."
					return nil, nil, message, http.StatusConflict
				}
			} else {
				message := "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{4,}$."
				return nil, nil, message, http.StatusUnprocessableEntity
			}
		}
	}

	return nil, nil, "", 200
}

func (service *UrlService) GetShortCode(shortcode string) (*httpEntity.Url, error, string, int) {

	url, err, message, status := service.urlRepository.UpdateUrl(shortcode)
	result := httpEntity.Url{}
	copier.Copy(&result, url)
	return &result, err, message, status
}

func (service *UrlService) GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int) {
	var result httpEntity.ShortcodeStats
	url, err, message, status := service.urlRepository.GetStats(shortcode)
	if url != nil {
		result.LastSeenDate = url.LastSeenDate
		result.RedirectCount = url.RedirectCount
		result.StartDate = url.CreatedAt
		return &result, err, message, status
	}
	return nil, err, message, status
}
