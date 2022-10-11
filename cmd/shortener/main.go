package main

import (
	"github.com/vovanwin/shorter"
	"github.com/vovanwin/shorter/internal/app/config"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s := shorter.CreateNewServer()
	s.MountHandlers()
	http.ListenAndServe(config.Domain, s.Router)
}
