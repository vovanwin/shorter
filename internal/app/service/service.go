package service

import (
	"github.com/vovanwin/shorter/internal/app/model"
	"github.com/vovanwin/shorter/internal/app/repository"
)

type LinkService interface {
	GetLink(code string) (model.URLLink, error)
	AddLink(model model.URLLink) error
}

type Service struct {
	LinkService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{LinkService: NewFile(repo.LinkService)}
}
