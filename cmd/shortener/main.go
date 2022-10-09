package main

import (
	"github.com/vovanwin/shorter"
	"github.com/vovanwin/shorter/internal/app/config"
	"net/http"
)

func main() {
	s := shorter.CreateNewServer()
	s.MountHandlers()
	http.ListenAndServe(config.Domain, s.Router)
}
