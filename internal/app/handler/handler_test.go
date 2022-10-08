package handler

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"time"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {

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
			name: "Есть ссылка в body",
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
		t.Run(tt.name, func(t *testing.T) {
			if tt.want.method == http.MethodPost {
				bodyReader := strings.NewReader(tt.want.body)
				request := httptest.NewRequest(tt.want.method, tt.want.path, bodyReader)

				// создаём новый Recorder
				w := httptest.NewRecorder()
				// определяем хендлер
				h := http.HandlerFunc(Run)
				// запускаем сервер
				h.ServeHTTP(w, request)
				res := w.Result()
				_ = res.Body.Close()
				assert.Equal(t, tt.want.code, res.StatusCode)
			}

			if tt.want.method == http.MethodGet {
				var newURL = urlLink{
					ID:    time.Now().UnixNano(),
					Long:  tt.want.body,
					Short: tt.want.path,
				}
				array = append(array, newURL)

				request := httptest.NewRequest(tt.want.method, "/"+tt.want.path, nil)

				// создаём новый Recorder
				w := httptest.NewRecorder()
				// определяем хендлер
				h := http.HandlerFunc(Run)
				// запускаем сервер
				h.ServeHTTP(w, request)
				res := w.Result()

				_ = res.Body.Close()

				assert.Equal(t, tt.want.code, res.StatusCode)
			}
		})
	}
}
