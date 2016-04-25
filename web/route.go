package web

import "github.com/senghoo/captain/web/controllers"

func (s *Server) initRoute() {
	s.m.Get("/", controllers.Main)
}
