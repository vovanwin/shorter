package repository

import (
	"errors"
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

func (m *Memory) AddLink(model model.URLLink) error {
	mu.Lock()
	defer mu.Unlock()

	array = append(array, model)
	return nil
}
