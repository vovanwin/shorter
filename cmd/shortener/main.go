package main

import (
	"flag"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/handler"
	"github.com/vovanwin/shorter/internal/app/repository"
	"github.com/vovanwin/shorter/internal/app/service"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	var repositoryhandler repository.LinkService
	conf := new(config.Config)

	fileStoragePath := conf.GetConfig().FileStoragePath
	if fileStoragePath == "" {
		repositoryhandler = repository.NewMemory()
	} else {
		repositoryhandler = repository.NewJSON()
	}
	repos := repository.NewRepository(repositoryhandler)
	services := service.NewService(repos)

	s := handler.CreateNewServer(services)
	s.MountHandlers()

	http.ListenAndServe(s.Config.GetConfig().ServerAddress, s.Router)
}
