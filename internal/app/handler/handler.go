package handler

import (
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/helper"
	"io"
	"net/http"
	"time"
)

var array []urlLink

type urlLink struct {
	ID    int64  `json:"ID"`
	Long  string `json:"Long"`
	Short string `json:"Short"`
}

// Run Обработчик.
func Run(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHandler(w, r)
	}

	if r.Method == http.MethodPost {
		postHandler(w, r)
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	for _, value := range array {
		if value.Short == path {
			w.Header().Set("Location", value.Long)
			w.WriteHeader(http.StatusTemporaryRedirect)
			w.Write([]byte(value.Long))
		}
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	longLink := string(data[:])
	u := helper.IsURL(longLink)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if u == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code := helper.NewCode()
	var newURL = urlLink{
		ID:    time.Now().UnixNano(),
		Long:  longLink,
		Short: code,
	}
	array = append(array, newURL)

	w.WriteHeader(http.StatusCreated)
	shortLink := helper.Concat2builder("http://", config.Domain, "/", code)
	w.Write([]byte(shortLink))
}
