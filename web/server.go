package web

import (
	"path"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/session"
	"github.com/senghoo/captain/modules/settings"
	"github.com/senghoo/captain/web/middleware"
	"gopkg.in/macaron.v1"
)

type Server struct {
	m *macaron.Macaron
}

func NewServer() *Server {
	return &Server{
		m: macaron.Classic(),
	}
}

func (s *Server) Run() {
	s.initRoute()
	s.initMiddleWare()
	s.m.SetAutoHead(true)
	s.m.Run()
}

func (s *Server) initMiddleWare() {
	s.m.Use(macaron.Logger())
	s.m.Use(pongo2.Pongoer(pongo2.Options{
		Directory: path.Join(settings.GetStaticPath(), "templates"),
	}))
	s.m.Use(macaron.Recovery())
	s.m.Use(macaron.Static(path.Join(settings.GetStaticPath(), "public")))
	s.m.Use(macaron.Static(path.Join(settings.GetStaticPath(), "assets")))
	s.m.Use(cache.Cacher())
	s.m.Use(session.Sessioner())
	s.m.Use(csrf.Csrfer(csrf.Options{
		Secret:    settings.GetOrDefault("csrf.key", "development csrf keys"),
		SetCookie: true,
		Header:    "X-Csrf-Token",
	}))
	s.m.Use(middleware.Contexter())
}
