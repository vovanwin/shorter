package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/vovanwin/shorter/internal/app/model"
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
		h := CreateNewServer()
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
			ID:    time.Now().UnixNano(),
			Long:  tt.want.body,
			Short: tt.want.path,
		}
		array = append(array, newURL)

		request := httptest.NewRequest(tt.want.method, "/"+tt.want.path, nil)

		w := httptest.NewRecorder()
		h := CreateNewServer()
		h.MountHandlers()
		h.Router.ServeHTTP(w, request)
		res := w.Result()

		_ = res.Body.Close()

		assert.Equal(t, tt.want.code, res.StatusCode)
	}
}

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {

	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	s.Router.Get("/{shortUrl}", Redirect)
	s.Router.Post("/", CreateShortLink)
}
