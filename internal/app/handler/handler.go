package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vovanwin/shorter/internal/app/helper"
	"github.com/vovanwin/shorter/internal/app/middleware"
	"github.com/vovanwin/shorter/internal/app/model"
	"io"
	"net/http"
	"time"
)

func (s *Server) Redirect(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "shortUrl")

	url, err := s.Service.GetLink(path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url.Long)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (s *Server) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.Key).([]byte)

	var ret [16]byte
	copy(ret[:], user)

	reader, err := helper.ReadRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := io.ReadAll(reader)
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
	shortLink := helper.Concat2builder(s.Config.GetConfig().ServerAddress, "/", code)

	var newURL = model.URLLink{ID: time.Now().UnixNano(), Long: longLink, ShortLink: shortLink, Code: code, UserID: ret}

	err = s.Service.AddLink(newURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortLink))
}

func (s *Server) ShortHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.Key).([]byte)

	reader, err := helper.ReadRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	code := helper.NewCode()

	var ret [16]byte
	copy(ret[:], user)

	var newURL = model.URLLink{ID: time.Now().UnixNano(), Code: code, UserID: ret}

	err = json.NewDecoder(reader).Decode(&newURL)
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
	shortLink := helper.Concat2builder(s.Config.GetConfig().ServerAddress, "/", code)
	newURL.ShortLink = shortLink

	err = s.Service.AddLink(newURL)
	if err != nil {
		fmt.Println(err)
	}

	var ReturnURL = model.URLLink{ShortLink: newURL.ShortLink, UserID: newURL.UserID}

	res, err := json.Marshal(ReturnURL)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (s *Server) GetUserURL(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.Key).([]byte)
	var ret [16]byte
	copy(ret[:], user)
	urls, err := s.Service.GetLinksUser(ret)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := json.Marshal(urls)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (s *Server) BatchShorten(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.Key).([]byte)
	reader, err := helper.ReadRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var ret [16]byte
	copy(ret[:], user)

	type UserURLLinks struct {
		Correlation string `json:"correlation_id,omitempty"`
		OriginalUrl string `json:"original_url,omitempty"`
	}

	type UserURLLinksResponse struct {
		Correlation string `json:"correlation_id,omitempty"`
		ShortUrl    string `json:"short_url,omitempty"`
	}

	var arrURL []UserURLLinks

	//err = json.NewDecoder(reader).Decode(&arrURL)
	err = json.NewDecoder(reader).Decode(&arrURL)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newURL = model.URLLink{}
	var arrURLResponse []UserURLLinksResponse

	for _, urlValue := range arrURL {
		u := helper.IsURL(urlValue.OriginalUrl)
		if !u {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		code := helper.NewCode()
		newURL.ID = time.Now().UnixNano()
		newURL.UserID = ret
		newURL.Code = code
		shortLink := helper.Concat2builder(s.Config.GetConfig().ServerAddress, "/", code)

		newURL.ShortLink = shortLink
		newURL.Long = urlValue.OriginalUrl
		err = s.Service.AddLink(newURL)
		if err != nil {
			fmt.Println(err)
		}
		arrURLResponse = append(arrURLResponse, UserURLLinksResponse{ShortUrl: shortLink, Correlation: urlValue.Correlation})
	}

	res, err := json.Marshal(arrURLResponse)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {

	err := s.Service.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
