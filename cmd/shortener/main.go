package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/vovanwin/shorter"
	"github.com/vovanwin/shorter/internal/app/config"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	rand.Seed(time.Now().UnixNano())
	s := shorter.CreateNewServer()
	s.MountHandlers()
	http.ListenAndServe(cfg.SERVER_ADDRESS, s.Router)
}
