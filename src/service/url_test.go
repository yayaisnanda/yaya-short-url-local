package services

import (
	"testing"
	"time"
	dbEntity "yaya_short_url_local/src/entity/db"

	// modelAPI "example_app/entity/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// SETUP
type repositoryDBMock struct {
	mock.Mock
}

func (repository *repositoryDBMock) ShortenUrl(urlData *dbEntity.Url) (*dbEntity.Url, error, string, int) {
	repository.Called(urlData)
	urlData.ID = 1
	urlData.Url = "test url"
	urlData.UrlShort = "test shorturl"
	return nil, nil, "", 0
}

func (repository *repositoryDBMock) GetShortCode(shortcode string) (*dbEntity.Url, error, string, int) {
	repository.Called(shortcode)
	shortcode = "test shortcode"
	return nil, nil, "", 0
}

func (repository *repositoryDBMock) GetStats(shortcode string) (*dbEntity.Url, error, string, int) {
	repository.Called(shortcode)
	shortcode = "test shortcode"
	return nil, nil, "", 0
}

func TestGETSHORTCODEMocked(t *testing.T) {
	t.Parallel()
	now := time.Now()
	dbMockData := repositoryDBMock{}
	var shortcode string = "jklmn"
	dbMockData.On("UpdateUrlByID", shortcode, &dbEntity.Url{
		UpdatedAt:     now,
		LastSeenDate:  &now,
		RedirectCount: 1,
	}).Return(nil)
	urlService := UrlService{&dbMockData}
	resultFuncService, _, _, _ := urlService.GetShortCode(shortcode)
	// {
	// 	Name:         "Test Update",
	// 	IDCardNumber: "IDCARDUPDATE1213243",
	// 	Address:      "Adress Update 96",
	// }
	assert.Equal(t, resultFuncService, nil, "It should be true")
}
