package model

import "github.com/google/uuid"

type URLLink struct {
	ID        int64     `json:"-"`
	Long      string    `json:"url,omitempty"`
	Code      string    `json:"code,omitempty"`
	ShortLink string    `json:"result,omitempty"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
}

type UserURLLinks struct {
	Long      string `json:"original_url,omitempty"`
	ShortLink string `json:"short_url,omitempty"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Long      string    `json:"url,omitempty"`
	Code      string    `json:"code,omitempty"`
	ShortLink string    `json:"result,omitempty"`
}
