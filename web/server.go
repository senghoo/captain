package web

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/session"
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
	s.m.SetAutoHead(true)
	s.m.Run()
}

func (s *Server) initMiddleWare() {
	s.m.Use(macaron.Logger())
	s.m.Use(macaron.Recovery())
	s.m.Use(macaron.Static("public"))
	s.m.Use(macaron.Static("assets"))
	s.m.Use(cache.Cacher())
	s.m.Use(session.Sessioner())
	s.m.Use(pongo2.Pongoer())
}
