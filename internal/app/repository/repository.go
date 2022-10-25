package repository

import (
	"github.com/vovanwin/shorter/internal/app/model"
	"sync"
)

var mu sync.Mutex // Объявляет мьютекс
var array []model.URLLink

type LinkService interface {
	GetLink(code string) (model.URLLink, error)
	AddLink(model model.URLLink) error
}

type Repository struct {
	LinkService
}

func NewRepository(repo LinkService) *Repository {
	return &Repository{
		LinkService: repo,
	}
}
