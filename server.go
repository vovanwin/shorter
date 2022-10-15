package shorter

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vovanwin/shorter/internal/app/handler"
)

type Server struct {
	Router  *chi.Mux
	Handler handler.Handler
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {

	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	s.Router.Post("/api/shorten", s.Handler.ShortHandler)
	s.Router.Get("/{shortUrl}", s.Handler.Redirect)
	s.Router.Post("/", s.Handler.CreateShortLink)
}
