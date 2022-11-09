package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vovanwin/shorter/internal/app/config"
	customMiddleware "github.com/vovanwin/shorter/internal/app/middleware"
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

	s.Router.Use(customMiddleware.UserCookie)
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.AllowContentEncoding("gzip"))
	s.Router.Use(middleware.Compress(5, "application/json", "text/plain", "application/x-gzip"))

	s.Router.Get("/api/user/urls", s.GetUserURL)
	s.Router.Post("/api/shorten", s.ShortHandler)
	s.Router.Get("/ping", s.Ping)
	s.Router.Get("/{shortUrl}", s.Redirect)
	s.Router.Post("/", s.CreateShortLink)
}
