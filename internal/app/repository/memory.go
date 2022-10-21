package repository

import (
	"errors"
	"github.com/vovanwin/shorter/internal/app/model"
)

type Memory struct {
}

var array []model.URLLink

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) GetLink(code string) (model.URLLink, error) {
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
	array = append(array, model)
	return nil
}
