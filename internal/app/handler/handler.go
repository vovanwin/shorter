package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/helper"
	"github.com/vovanwin/shorter/internal/app/model"
	"io"
	"net/http"
	"time"
)

var array []model.URLLink

type Handler struct {
	config config.Config
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "shortUrl")
	for _, value := range array {
		if value.Code == path {
			w.Header().Set("Location", value.Long)
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	longLink := string(data[:])
	u := helper.IsURL(longLink)

	if !u {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code := helper.NewCode()
	var newURL = model.URLLink{ID: time.Now().UnixNano(), Long: longLink, Code: code}
	array = append(array, newURL)

	w.WriteHeader(http.StatusCreated)
	shortLink := helper.Concat2builder("http://", h.config.SERVER_ADDRESS, "/", code)
	w.Write([]byte(shortLink))
}

func (h *Handler) ShortHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	code := helper.NewCode()
	var newURL = model.URLLink{ID: time.Now().UnixNano(), Code: code}

	err := json.NewDecoder(r.Body).Decode(&newURL)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := helper.IsURL(newURL.Long)

	if !u {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shortLink := helper.Concat2builder("http://", h.config.SERVER_ADDRESS, "/", code)
	newURL.ShortLink = shortLink

	array = append(array, newURL)

	w.WriteHeader(http.StatusCreated)
	var ReturmURL = model.URLLink{ShortLink: newURL.ShortLink}

	res, err := json.Marshal(ReturmURL)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(res)
}
