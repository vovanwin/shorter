package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/vovanwin/shorter/internal/app/model"
)

type Memory struct {
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) GetLink(code string) (model.URLLink, error) {
	mu.Lock()
	defer mu.Unlock()

	var data model.URLLink
	var err error
	for _, value := range array {
		if value.Code == code {
			data = value
			break
		}
	}

	if (model.URLLink{} == data) {
		err = errors.New("ссылка не найдена")
	}
	return data, err
}

func (m *Memory) GetLinkByLong(code string) (model.URLLink, error) {
	mu.Lock()
	defer mu.Unlock()

	var data model.URLLink
	var err error
	for _, value := range array {
		if value.Code == code {
			data = value
			break
		}
	}

	if (model.URLLink{} == data) {
		err = errors.New("ссылка не найдена")
	}
	return data, err
}

func (m *Memory) AddLink(model model.URLLink) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	array = append(array, model)
	return "", nil
}

func (m *Memory) GetLinksUser(user uuid.UUID) ([]model.UserURLLinks, error) {
	var response []model.UserURLLinks
	var data model.UserURLLinks
	var err error

	for _, value := range array {
		if value.UserID == user {
			data = model.UserURLLinks{ShortLink: value.ShortLink, Long: value.Long}
			response = append(response, data)
		}
	}

	if (model.UserURLLinks{} == data) {
		err = errors.New("ссылка не найдена")
		return nil, err
	}
	return response, nil
}

func (m *Memory) Ping() error {
	var err = errors.New("репозиторий не поддерживает БД")
	return err
}
