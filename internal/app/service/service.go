package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vovanwin/shorter/internal/app/model"
	"github.com/vovanwin/shorter/internal/app/repository"
)

type LinkService interface {
	GetLink(code string) (model.URLLink, error)
	GetLinkByLong(long string) (model.URLLink, error)
	GetLinksUser(user uuid.UUID) ([]model.UserURLLinks, error)
	AddLink(model model.URLLink) (string, error)
	Ping() error
}

type Service struct {
	LinkService
}

func NewService(repo *repository.Repository) *Service {
	fmt.Printf("%s", repo)
	return &Service{LinkService: NewFile(repo.LinkService)}
}
