package http

import "time"

type UrlForm struct {
	URL       string `json:"url"`
	Shortcode string `json:"shortcode"`
}

type Url struct {
	ID        uint      `json:"id"`
	Url       string    `json:"url"`
	UrlShort  string    `json:"shorturl"`
	CreatedAt time.Time `json:"created_at"`
}

type ShortcodeStats struct {
	StartDate     time.Time  `json:"start_date"`
	LastSeenDate  *time.Time `json:"last_seen_date,omitempty"`
	RedirectCount int        `json:"redirect_count"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
