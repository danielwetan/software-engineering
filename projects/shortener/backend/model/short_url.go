package model

type ShortUrl struct {
	ShortCode string `json:"shortcode"`
	UserID    int    `json:"user_id"`
	Target    string `json:"target"`
	CreatedAt string `json:"created_at"`
}

// Request
type (
	CreateShortUrlRequest struct {
		UserID    int    `json:"user_id,omitempty"`
		Target    string `json:"target" validate:"required"`
		ShortCode string
	}

	GetShortUrlRequest struct {
		ShortCode string `json:"shortcode" validate:"required"`
	}
)

// Response
type (
	CreateShortUrlResponse struct {
		ShortCode string `json:"shortcode"`
	}
)
