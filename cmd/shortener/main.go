package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length     = len(alphabet)
	CodeLength = 6
	Domain     = "127.0.0.1:8080"
)

var array []urlLink

// Обработчик.
func run(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		path := r.URL.Path[1:]

		for _, value := range array {
			if value.Short == path {
				w.Header().Set("Location", value.Long)
				w.WriteHeader(http.StatusTemporaryRedirect)
				w.Write([]byte(value.Long))
			}
		}
	}

	if r.Method == http.MethodPost {
		data, err := io.ReadAll(r.Body)
		longLink := string(data[:])
		u := IsURL(longLink)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if u == false {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		println(u)

		code := NewCode()
		var newURL = urlLink{
			ID:    time.Now().UnixNano(),
			Long:  longLink,
			Short: code,
		}
		array = append(array, newURL)

		w.WriteHeader(http.StatusCreated)
		shortLink := concat2builder("http://", Domain, "/", code)
		w.Write([]byte(shortLink))
	}

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", run)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(Domain, mux)
	log.Fatal(err)
}

type urlLink struct {
	ID    int64  `json:"ID"`
	Long  string `json:"Long"`
	Short string `json:"Short"`
}

func NewCode() string {
	var letters = []rune(alphabet)

	code := make([]rune, CodeLength)
	for i := range code {
		code[i] = letters[rand.Intn(length)]
	}

	return string(code)
}

func concat2builder(http, x, z, y string) string {
	var builder strings.Builder
	builder.Grow(len(http) + len(x) + len(z) + len(y)) // Только эта строка выделяет память
	builder.WriteString(http)
	builder.WriteString(x)
	builder.WriteString(z)
	builder.WriteString(y)
	return builder.String()
}

func IsURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
