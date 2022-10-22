package service

import (
	"github.com/vovanwin/shorter/internal/app/model"
	"github.com/vovanwin/shorter/internal/app/repository"
)

type Link struct {
	repo repository.LinkService
}

func NewFile(repo repository.LinkService) *Link {
	return &Link{repo: repo}
}

func (s *Link) GetLink(code string) (model.URLLink, error) {
	return s.repo.GetLink(code)
}

func (s *Link) AddLink(model model.URLLink) error {
	return s.repo.AddLink(model)
}
