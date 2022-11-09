package main

import (
	"context"
	"flag"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/handler"
	"github.com/vovanwin/shorter/internal/app/repository"
	"github.com/vovanwin/shorter/internal/app/service"
	"github.com/vovanwin/shorter/pkg/client/postgresql"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	var repositoryhandler repository.LinkService
	conf := new(config.Config)

	pgConfig := postgresql.NewPgConfig(
		conf.GetConfig().DatabaseDsn,
	)

	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		fileStoragePath := conf.GetConfig().FileStoragePath
		if fileStoragePath == "" {
			repositoryhandler = repository.NewMemory()
		} else {
			repositoryhandler = repository.NewJSON()
		}
	} else {
		repositoryhandler = repository.NewDB(pgClient)
	}

	repos := repository.NewRepository(repositoryhandler)
	services := service.NewService(repos)

	s := handler.CreateNewServer(services)
	s.MountHandlers()

	http.ListenAndServe(s.Config.GetConfig().ServerAddress, s.Router)
}
