package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/helper"
	"github.com/vovanwin/shorter/internal/app/model"
	"io"
	"log"
	"net/http"
	"time"
)

var array []model.UrlLink

func Redirect(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "shortUrl")
	for _, value := range array {
		if value.Short == path {
			w.Header().Set("Location", value.Long)
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func CreateShortLink(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal("Ошибка чтения тела запроса")
		return
	}

	longLink := string(data[:])
	u := helper.IsURL(longLink)

	if !u {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Ошибка проверки ссылки на валидность")
		return
	}

	code := helper.NewCode()
	var newURL = model.UrlLink{ID: time.Now().UnixNano(), Long: longLink, Short: code}
	array = append(array, newURL)

	w.WriteHeader(http.StatusCreated)
	shortLink := helper.Concat2builder("http://", config.Domain, "/", code)
	w.Write([]byte(shortLink))
}
