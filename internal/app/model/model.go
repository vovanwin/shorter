package model

type URLLink struct {
	ID        int64  `json:"-"`
	Long      string `json:"url,omitempty"`
	Code      string `json:"code,omitempty"`
	ShortLink string `json:"result,omitempty"`
}
