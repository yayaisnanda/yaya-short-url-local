package db

import (
	"database/sql"
	"net/http"
	"time"
	dbEntity "yaya_short_url_local/src/entity/db"
	connection "yaya_short_url_local/src/util/helper/mysqlconnection"

	"github.com/jinzhu/gorm"
)

type UrlRepository struct {
	DB gorm.DB
}

func UrlRepositoryHandler() *UrlRepository {
	return &UrlRepository{DB: *connection.GetConnection()}
}

type UrlRepositoryInterface interface {
	// ShortenUrl(urlData *dbEntity.Url) (*dbEntity.Url, error, string, int)
	// GetShortCode(shortcode string) (*dbEntity.Url, error, string, int)
	// GetStats(shortcode string) (*dbEntity.Url, error, string, int)

	GetUrlList() ([]*dbEntity.Url, error, string, int)
	InsertUrl(urlData *dbEntity.Url) (error, string, int)
	UpdateUrl(shortcode string) (*dbEntity.Url, error, string, int)
	GetStats(shortcode string) (*dbEntity.Url, error, string, int)
}

// func (repository *UrlRepository) ShortenUrl(urlData *dbEntity.Url) (*dbEntity.Url, error, string, int) {
// 	//	url := dbEntity.Url{}

// 	var urls []dbEntity.Url

// 	err := repository.DB.Where("deleted_at IS NULL").Find(&urls)
// 	if err.Error != nil {
// 		message := "Failed get url list"
// 		return nil, err.Error, message, http.StatusBadRequest
// 	}

// 	if len(urls) == 0 {
// 		if urlData.UrlShort == "" {
// 			g, err := reggen.NewGenerator("^[0-9a-zA-Z_]{6}$")
// 			if err != nil {
// 				panic(err)
// 			}
// 			generate := g.Generate(0)
// 			urlData.UrlShort = generate

// 			error := repository.DB.Create(&urlData)
// 			if error.Error != nil {
// 				message := "Failed to insert url"
// 				return nil, error.Error, message, http.StatusBadRequest
// 			}
// 			message := "Created"
// 			return nil, nil, message, http.StatusCreated
// 		} else {
// 			r := regexp.MustCompile(`^[0-9a-zA-Z_]{4,}$`)
// 			matches := r.FindAllString(urlData.UrlShort, 1)

// 			checkRegex := false
// 			if len(matches) > 0 {
// 				checkRegex = true
// 			}

// 			if checkRegex {
// 				//fmt.Println("if")
// 				urlData.UrlShort = matches[0]

// 				error := repository.DB.Create(&urlData)
// 				if error.Error != nil {
// 					message := "Failed to insert url"
// 					return nil, error.Error, message, http.StatusBadRequest
// 				}
// 				message := "Created"
// 				return nil, nil, message, http.StatusCreated
// 			} else {
// 				//not pass regex
// 				//fmt.Println("else")
// 				message := "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{4,}$."
// 				return nil, err.Error, message, http.StatusUnprocessableEntity
// 			}

// 		}
// 	} else {
// 		if urlData.UrlShort == "" {
// 			for {
// 				g, err := reggen.NewGenerator("^[0-9a-zA-Z_]{6}$")
// 				if err != nil {
// 					panic(err)
// 				}
// 				generate := g.Generate(0)
// 				// fmt.Println(generate)
// 				// check with data on DB insert db
// 				check := false
// 				for i, _ := range urls {
// 					if urls[i].UrlShort == generate {
// 						check = true
// 					}
// 				}
// 				if !check {
// 					// inssert shorten code to database
// 					urlData.UrlShort = generate

// 					error := repository.DB.Create(&urlData)
// 					if error.Error != nil {
// 						message := "Failed to insert url"
// 						return nil, error.Error, message, http.StatusBadRequest
// 					}
// 					message := "Created"
// 					return nil, nil, message, http.StatusCreated
// 				}
// 			}
// 		} else {
// 			r := regexp.MustCompile(`^[0-9a-zA-Z_]{4,}$`)
// 			matches := r.FindAllString(urlData.UrlShort, 1)
// 			checkRegex := false
// 			if len(matches) > 0 {
// 				checkRegex = true
// 			}
// 			if checkRegex {
// 				//check with data on DB insert database
// 				check := false
// 				for i, _ := range urls {
// 					if urls[i].UrlShort == urlData.UrlShort {
// 						check = true
// 					}
// 				}
// 				if !check {
// 					// fmt.Println("insert to db")
// 					// inssert shorten code to database
// 					urlData.UrlShort = matches[0]

// 					error := repository.DB.Create(&urlData)
// 					if error.Error != nil {
// 						message := "Failed to insert url"
// 						return nil, error.Error, message, http.StatusBadRequest
// 					}

// 					message := "Created"
// 					return nil, nil, message, http.StatusCreated
// 				} else {
// 					// shorcode exist in database
// 					message := "The the desired shortcode is already in use. Shortcodes are case-sensitive."
// 					return nil, err.Error, message, http.StatusConflict
// 				}
// 			} else {
// 				//not pass regex
// 				message := "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{4,}$."
// 				return nil, err.Error, message, http.StatusUnprocessableEntity
// 				// fmt.Println("not pass regex")
// 			}
// 		}
// 	}

// 	// query := repository.DB.Table("urls")
// 	// query = query.Create(&urlData)
// 	return nil, nil, "", http.StatusOK
// }

// func (repository *UrlRepository) GetShortCode(shortcode string) (*dbEntity.Url, error, string, int) {
// 	var url dbEntity.Url
// 	var newUrl dbEntity.Url

// 	now := time.Now()

// 	err := repository.DB.Where("deleted_at IS NULL AND url_short = ?", shortcode).Find(&url)

// 	if err.Error != nil || sql.ErrNoRows == nil {
// 		//common.Error(err.Error, "Failed get url from database: ")

// 		message := "The shortcode cannot be found in the system"
// 		return nil, err.Error, message, http.StatusNotFound
// 	}

// 	// lastSeenDate := time.Now().UTC().Format("2006-01-02T15:04:05.000+0000")
// 	// ts, errr := time.Parse("2006-01-02T15:04:05.000+0000", lastSeenDate)
// 	// if errr != nil {
// 	// 	fmt.Println(errr)
// 	// 	return nil, errr, "", http.StatusForbidden
// 	// }
// 	//ts.Format(time.RFC3339)

// 	//fmt.Println(ts)

// 	newUrl.UpdatedAt = now
// 	newUrl.LastSeenDate = &now
// 	newUrl.RedirectCount = url.RedirectCount + 1
// 	err = repository.DB.Model(&url).Update(newUrl)

// 	if err.Error != nil {
// 		message := "can't update data to database"
// 		return nil, err.Error, message, http.StatusBadGateway
// 	}
// 	return &url, err.Error, "", http.StatusFound

// }

// func (repository *UrlRepository) GetStats(shortcode string) (*dbEntity.Url, error, string, int) {
// 	var url dbEntity.Url

// 	err := repository.DB.Where("deleted_at IS NULL AND url_short = ?", shortcode).Find(&url)

// 	if err.Error != nil || sql.ErrNoRows == nil {
// 		//common.Error(err.Error, "Failed get url from database: ")
// 		message := "The shortcode cannot be found in the system"
// 		return nil, err.Error, message, http.StatusNotFound
// 	}

// 	// lastSeenDate := url.LastSeenDate.Format("2006-01-02T15:04:05.000+0000")

// 	// // t, errr := time.Parse("2006-01-02T15:04:05-0700", lastSeenDate)
// 	// // if errr != nil {
// 	// // 	fmt.Println(errr)
// 	// // 	return nil, errr, "", http.StatusForbidden
// 	// // }
// 	// ts, errr := time.Parse("2006-01-02T15:04:05.000+0000", lastSeenDate)
// 	// if errr != nil {
// 	// 	fmt.Println(errr)
// 	// 	return nil, errr, "", http.StatusForbidden
// 	// }

// 	url.LastSeenDate = url.LastSeenDate

// 	//fmt.Println(url.LastSeenDate)

// 	return &url, err.Error, "success get shorcode stats", http.StatusFound

// }

func (repository *UrlRepository) GetUrlList() ([]*dbEntity.Url, error, string, int) {
	var urls []*dbEntity.Url

	err := repository.DB.Where("deleted_at IS NULL").Find(&urls)
	if err.Error != nil {
		message := "Failed get url list"
		return nil, err.Error, message, http.StatusBadRequest
	}
	message := "Success get url list"
	return urls, err.Error, message, http.StatusOK
}

func (repository *UrlRepository) InsertUrl(urlData *dbEntity.Url) (error, string, int) {
	error := repository.DB.Create(&urlData)
	if error.Error != nil {
		message := "Failed to insert url"
		return error.Error, message, http.StatusBadRequest
	}
	message := "Created"
	return nil, message, http.StatusCreated
}

func (repository *UrlRepository) UpdateUrl(shortcode string) (*dbEntity.Url, error, string, int) {
	var url dbEntity.Url
	var newUrl dbEntity.Url

	now := time.Now()

	err := repository.DB.Where("deleted_at IS NULL AND url_short = ?", shortcode).Find(&url)

	if err.Error != nil || sql.ErrNoRows == nil {
		message := "The shortcode cannot be found in the system"
		return nil, err.Error, message, http.StatusNotFound
	}
	newUrl.UpdatedAt = now
	newUrl.LastSeenDate = &now
	newUrl.RedirectCount = url.RedirectCount + 1
	err = repository.DB.Model(&url).Update(newUrl)

	if err.Error != nil {
		message := "can't update data to database"
		return nil, err.Error, message, http.StatusBadGateway
	}
	return &url, err.Error, "", http.StatusFound
}

func (repository *UrlRepository) GetStats(shortcode string) (*dbEntity.Url, error, string, int) {
	var url dbEntity.Url

	err := repository.DB.Where("deleted_at IS NULL AND url_short = ?", shortcode).Find(&url)

	if err.Error != nil || sql.ErrNoRows == nil {
		message := "The shortcode cannot be found in the system"
		return nil, err.Error, message, http.StatusNotFound
	}

	url.LastSeenDate = url.LastSeenDate
	return &url, err.Error, "success get shorcode stats", http.StatusFound
}
