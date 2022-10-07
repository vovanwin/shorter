package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length     = len(alphabet)
	CodeLength = 6
	Domain     = "localhost:8080"
)

var array []url

// Обработчик.
func run(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		path := r.URL.Path[1:]

		var newUrl = url{
			ID:    time.Now().UnixNano(),
			Long:  "sdfsd",
			Short: "dsfsdf",
		}
		array = append(array, newUrl)

		j, err := json.Marshal(array)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(j))
		}
		for _, value := range array {
			if value.Short == path {
				fmt.Println(value.Short)
				w.WriteHeader(http.StatusTemporaryRedirect)
				w.Header().Set("Location", value.Long)
			}
		}
	}

	if r.Method == http.MethodPost {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		longLink := string(data[:])

		code := NewCode()
		var newUrl = url{
			ID:    time.Now().UnixNano(),
			Long:  longLink,
			Short: code,
		}
		array = append(array, newUrl)

		fmt.Println(longLink)
		fmt.Println(array)
		w.WriteHeader(http.StatusCreated)
		shortLink := concat2builder(Domain, "/", code)
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

type url struct {
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

func concat2builder(x, z, y string) string {
	var builder strings.Builder
	builder.Grow(len(x) + len(z) + len(y)) // Только эта строка выделяет память
	builder.WriteString(x)
	builder.WriteString(z)
	builder.WriteString(y)
	return builder.String()
}
