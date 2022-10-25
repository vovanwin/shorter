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

func NewJSON() *JSON {
	return &JSON{}
}

func init() {
	conf := new(JSON)

	jsonRead, err := helper.NewConsumer(conf.getPath())
	if err != nil {
		return
	}

	urls, _ := jsonRead.ReadEvent()
	jsonRead.Close()
	array = urls
}

func (j *JSON) GetLink(code string) (model.URLLink, error) {
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

func (j *JSON) AddLink(model model.URLLink) error {
	jsonRead, err := helper.NewConsumer(j.getPath())
	if err != nil {
		return err
	}

	array = append(array, model)

	json, err := helper.NewProducer(j.getPath())
	if err != nil {
		return err
	}

	err = json.WriteEvent(&array)
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
