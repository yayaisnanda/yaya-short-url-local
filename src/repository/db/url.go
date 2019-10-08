package db

import (
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
	GetUrlList() ([]*dbEntity.Url, error)
	InsertUrl(urlData *dbEntity.Url) error
	UpdateUrl(shortcode string) (*dbEntity.Url, error)
	GetStats(shortcode string) (*dbEntity.Url, error)
}

func (repository *UrlRepository) GetUrlList() ([]*dbEntity.Url, error) {
	var urls []*dbEntity.Url
	err := repository.DB.Where("deleted_at IS NULL").Find(&urls)
	return urls, err.Error
}

func (repository *UrlRepository) InsertUrl(urlData *dbEntity.Url) error {
	error := repository.DB.Create(&urlData)
	return error.Error
}

func (repository *UrlRepository) UpdateUrl(shortcode string) (*dbEntity.Url, error) {
	var url dbEntity.Url
	var newUrl dbEntity.Url
	now := time.Now()
	err := repository.DB.Where("deleted_at IS NULL AND url_short = ?", shortcode).Find(&url)
	newUrl.UpdatedAt = now
	newUrl.LastSeenDate = &now
	newUrl.RedirectCount = url.RedirectCount + 1
	err = repository.DB.Model(&url).Update(newUrl)
	return &url, err.Error
}

func (repository *UrlRepository) GetStats(shortcode string) (*dbEntity.Url, error) {
	var res dbEntity.Url
	err := repository.DB.Where("deleted_at IS NULL AND url_short = ?", shortcode).Find(&res)
	return &res, err.Error
}

func (repository *UrlRepository) CheckUrlShort(shortcode string) (*dbEntity.Url, error) {
	var url dbEntity.Url
	err := repository.DB.Where("deleted_at IS NULL AND url_short = ?", shortcode).Find(&url)
	return &url, err.Error
}
