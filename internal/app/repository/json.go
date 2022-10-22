package repository

import (
	"errors"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/helper"
	"github.com/vovanwin/shorter/internal/app/model"
	"os"
	"strings"
)

type JSON struct {
	Config config.Config
}

var arrayURL []model.URLLink

func NewJSON() *JSON {
	return &JSON{}
}

func (j *JSON) GetLink(code string) (model.URLLink, error) {
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

func (j *JSON) AddLink(model model.URLLink) error {
	jsonRead, err := helper.NewConsumer(j.getPath())
	if err != nil {
		return err
	}

	urls, _ := jsonRead.ReadEvent()
	jsonRead.Close()

	arrayURL = append(urls, model)

	json, err := helper.NewProducer(j.getPath())
	if err != nil {
		return err
	}

	err = json.WriteEvent(&arrayURL)
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
func (j *JSON) getPath() string {
	test := strings.HasSuffix(os.Args[0], ".test")

	if test {
		return j.Config.GetConfig().FileStoragePathTest
	}
	return j.Config.GetConfig().FileStoragePath

}
