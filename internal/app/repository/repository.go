package repository

import (
	"github.com/google/uuid"
	"github.com/vovanwin/shorter/internal/app/model"
	"sync"
)

var mu sync.Mutex // Объявляет мьютекс
var array []model.URLLink

type LinkService interface {
	GetLink(code string) (model.URLLink, error)
	GetLinksUser(user uuid.UUID) ([]model.UserURLLinks, error)
	AddLink(model model.URLLink) error
	Ping() error
}

type Repository struct {
	LinkService
}

func NewRepository(repo LinkService) *Repository {
	return &Repository{
		LinkService: repo,
	}
}
