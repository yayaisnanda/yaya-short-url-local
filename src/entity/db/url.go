package db

import "time"

type Url struct {
	ID            uint       `gorm:"primary_key;column:id" json:"id"`
	Url           string     `gorm:"column:url" json:"url"`
	UrlShort      string     `gorm:"column:url_short" json:"shorturl"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	RedirectCount int        `gorm:"column:redirect_count" json:"redirect_count"`
	LastSeenDate  *time.Time `gorm:"column:last_seen_date" json:"last_seen_date"`
	DeletedAt     *time.Time `gorm:"deleted_at" json:"-"`
}

func (Url) TableName() string {
	return "urls"
}
