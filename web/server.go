package web

import "gopkg.in/macaron.v1"

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
	s.m.Run()
}
