package main

import (
	"log"
	"math/rand"
	"net/http"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length     = len(alphabet)
	CodeLength = 6
)

// Обработчик.
func run(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		encode := NewCode()

		w.Write([]byte(encode))
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Привет  "))
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", run)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func NewCode() string {
	var letters = []rune(alphabet)

	code := make([]rune, CodeLength)
	for i := range code {
		code[i] = letters[rand.Intn(length)]
	}

	return string(code)
}
