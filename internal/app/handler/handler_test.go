package handler

import (
	"github.com/stretchr/testify/assert"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/model"
	"github.com/vovanwin/shorter/internal/app/repository"
	"github.com/vovanwin/shorter/internal/app/service"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateShortLink(t *testing.T) {
	type want struct {
		code        int
		body        string
		contentType string
		method      string
		path        string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Нет ссылки в body",
			want: want{
				code:        400,
				contentType: "",
				method:      http.MethodPost,
				body:        "",
				path:        "/",
			},
		},
		{
			name: "Создание короткой ссылки",
			want: want{
				code:        201,
				contentType: "",
				method:      http.MethodPost,
				body:        "https://yandex.ru/search/?text=golang+%D0%BF%D1%80%D0%BE%D0%B2%D0%B5%D1%80%D0%B8%D1%82%D1%8C+%D1%82%D0%B8%D0%BF&lr=35",
				path:        "/",
			},
		},
		{
			name: "Не валидная ссылка",
			want: want{
				code:        400,
				contentType: "",
				method:      http.MethodPost,
				body:        "https://",
				path:        "/",
			},
		},
	}
	for _, tt := range tests {
		bodyReader := strings.NewReader(tt.want.body)
		request := httptest.NewRequest(tt.want.method, tt.want.path, bodyReader)

		w := httptest.NewRecorder()
		var repositoryhandler repository.LinkService
		conf := new(config.Config)

		fileStoragePath := conf.GetConfig().FileStoragePath
		if fileStoragePath == "" {
			repositoryhandler = repository.NewMemory()
		} else {
			repositoryhandler = repository.NewJSON()
		}
		repos := repository.NewRepository(repositoryhandler)
		services := service.NewService(repos)

		h := CreateNewServer(services)
		h.MountHandlers()
		h.Router.ServeHTTP(w, request)

		res := w.Result()
		_ = res.Body.Close()
		assert.Equal(t, tt.want.code, res.StatusCode)
	}
}

func TestRedirect(t *testing.T) {
	type want struct {
		code        int
		body        string
		contentType string
		method      string
		path        string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "редирект",
			want: want{
				code:        307,
				contentType: "Location",
				method:      http.MethodGet,
				body:        "https://yandex.ru/search/?text=golang+%D0%BF%D1%80%D0%BE%D0%B2%D0%B5%D1%80%D0%B8%D1%82%D1%8C+%D1%82%D0%B8%D0%BF&lr=35",
				path:        "iWvqZT",
			},
		},
	}

	for _, tt := range tests {
		var newURL = model.URLLink{
			ID:   time.Now().UnixNano(),
			Long: tt.want.body,
			Code: tt.want.path,
		}

		request := httptest.NewRequest(tt.want.method, "/"+tt.want.path, nil)

		w := httptest.NewRecorder()
		var repositoryhandler repository.LinkService
		conf := new(config.Config)

		fileStoragePath := conf.GetConfig().FileStoragePath
		if fileStoragePath == "" {
			repositoryhandler = repository.NewMemory()
		} else {
			repositoryhandler = repository.NewJSON()
		}
		repos := repository.NewRepository(repositoryhandler)
		services := service.NewService(repos)

		_, err := services.AddLink(newURL)
		if err != nil {
			return
		}

		h := CreateNewServer(services)
		h.MountHandlers()
		h.Router.ServeHTTP(w, request)
		res := w.Result()

		_ = res.Body.Close()

		assert.Equal(t, tt.want.code, res.StatusCode)
	}
}

func TestShortHandler(t *testing.T) {
	//TODO: Не знаю как прокинуть сюда конфиг кроме как этим способом
	conf := new(config.Config)

	type want struct {
		code        int
		body        string
		contentType string
		method      string
		path        string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Нет ссылки в body",
			want: want{
				code:        400,
				contentType: "",
				method:      http.MethodPost,
				body:        "",
				path:        conf.GetConfig().BaseURL,
			},
		},
		{
			name: "Создание короткой ссылки",
			want: want{
				code:        201,
				contentType: "",
				method:      http.MethodPost,
				body:        "{  \"url\" : \"https://yandex.ru/search/?text=golang+%D0%B4%D0%BE%D1%81%D1%82%D1%83%D1%82%D1%8C+%D0%B8%D0%B7+%D1%82%D0%B5%D0%BB%D0%B0+%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0&lr=35\"}",
				path:        conf.GetConfig().BaseURL,
			},
		},
		{
			name: "Не валидная ссылка",
			want: want{
				code:        400,
				contentType: "",
				method:      http.MethodPost,
				body:        "{  \"url\" : \"https://\"}",
				path:        conf.GetConfig().BaseURL,
			},
		},
	}
	for _, tt := range tests {
		bodyReader := strings.NewReader(tt.want.body)
		request := httptest.NewRequest(tt.want.method, tt.want.path, bodyReader)

		w := httptest.NewRecorder()

		rand.Seed(time.Now().UnixNano())

		var repositoryhandler repository.LinkService
		conf := new(config.Config)

		fileStoragePath := conf.GetConfig().FileStoragePath
		if fileStoragePath == "" {
			repositoryhandler = repository.NewMemory()
		} else {
			repositoryhandler = repository.NewJSON()
		}
		repos := repository.NewRepository(repositoryhandler)
		services := service.NewService(repos)

		h := CreateNewServer(services)
		h.MountHandlers()
		h.Router.ServeHTTP(w, request)

		res := w.Result()
		_ = res.Body.Close()
		assert.Equal(t, tt.want.code, res.StatusCode)
	}
}
