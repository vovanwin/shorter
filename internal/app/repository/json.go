package repository

import (
	"errors"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/helper"
	"github.com/vovanwin/shorter/internal/app/model"
	"os"
	"strings"
)

type Json struct {
	Config config.Config
}

var arrayUrl []model.URLLink

func NewJson() *Json {
	return &Json{}
}

func (j *Json) GetLink(code string) (model.URLLink, error) {
	var data model.URLLink
	jsonRead, err := helper.NewConsumer(j.getPath())
	urls, _ := jsonRead.ReadEvent()
	jsonRead.Close()

	for _, value := range urls {
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

func (j *Json) AddLink(model model.URLLink) error {
	jsonRead, err := helper.NewConsumer(j.getPath())
	if err != nil {
		return err
	}

	urls, _ := jsonRead.ReadEvent()
	jsonRead.Close()

	arrayUrl = append(urls, model)

	json, err := helper.NewProducer(j.getPath())
	if err != nil {
		return err
	}

	err = json.WriteEvent(&arrayUrl)
	if err != nil {
		return err
	}

	err = json.Close()
	if err != nil {
		return err
	}

	return nil
}

// Указывает путь в зависимости модульный тест это или реальный запуск приложения
func (j *Json) getPath() string {
	test := strings.HasSuffix(os.Args[0], ".test")

	if test {
		return j.Config.GetConfig().FileStoragePathTest
	}
	return j.Config.GetConfig().FileStoragePath

}
