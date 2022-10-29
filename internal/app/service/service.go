package service

import (
	"github.com/google/uuid"
	"github.com/vovanwin/shorter/internal/app/model"
	"github.com/vovanwin/shorter/internal/app/repository"
)

type LinkService interface {
	GetLink(code string) (model.URLLink, error)
	GetLinksUser(user uuid.UUID) ([]model.UserURLLinks, error)
	AddLink(model model.URLLink) error
}

type Service struct {
	LinkService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{LinkService: NewFile(repo.LinkService)}
}
