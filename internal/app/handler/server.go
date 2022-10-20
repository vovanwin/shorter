package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/service"
)

type Server struct {
	Service *service.Service
	Config  config.Config
	Router  *chi.Mux
}

func CreateNewServer(service *service.Service) *Server {
	s := &Server{Service: service}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {

	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	s.Router.Post("/api/shorten", s.ShortHandler)
	s.Router.Get("/{shortUrl}", s.Redirect)
	s.Router.Post("/", s.CreateShortLink)
}
