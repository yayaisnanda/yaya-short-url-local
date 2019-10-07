package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	httpEntity "yaya_short_url_local/src/entity/http"

	"github.com/stretchr/testify/assert"

	"testing"
)

type UrlServiceMock struct{}

// ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int)
// 	GetShortCode(shortcode string) (*httpEntity.Url, error, string, int)
// 	GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int)

func (service *UrlServiceMock) ShortenUrl(form httpEntity.UrlForm) (*httpEntity.Url, error, string, int) {
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	return &httpEntity.Url{
		ID:        1,
		Url:       "http:\\google.com",
		UrlShort:  "ktlo4",
		CreatedAt: t,
	}, nil, "", 0
}

func (service *UrlServiceMock) GetShortCode(shortcode string) (*httpEntity.Url, error, string, int) {
	return nil, nil, "", 0
}

func (service *UrlServiceMock) GetStats(shortcode string) (*httpEntity.ShortcodeStats, error, string, int) {
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	return &httpEntity.ShortcodeStats{
		StartDate:     t,
		LastSeenDate:  &t,
		RedirectCount: 1,
	}, nil, "", 0
}

func TestShortenUrlMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var idTest int = 1
	url := "/shorten"
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusFound, resp.Code, "Status should be 201")

	res := httpEntity.Url{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)

	assert.Equal(err, nil, "should have no error")
	assert.Equal(res.ID, idTest, "It should be same ID")
	assert.Equal(res.Url, "http:\\google.com", "It should be same URL")
	assert.Equal(res.UrlShort, "ktlo4", "It should be same ShortUrl")
}

func TestGetStatslMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var shortCodeTest string = "jklmn"
	fmt.Sprint(shortCodeTest)
	url := "/" + shortCodeTest + "/stats"
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusFound, resp.Code, "Status should be 201")

	res := httpEntity.ShortcodeStats{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)

	assert.Equal(err, nil, "should have no error")
	assert.Equal(res.StartDate, t, "It should be same t")
	assert.Equal(res.RedirectCount, 1, "It should be same RedirectCount")
	assert.Equal(res.LastSeenDate, t, "It should be same t")
}
