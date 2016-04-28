package web

import (
	"github.com/senghoo/captain/web/controllers"
	"github.com/senghoo/captain/web/controllers/user"
)

func (s *Server) initRoute() {
	s.m.Get("/", controllers.Main)
	s.m.Group("/user", func() {
		s.m.Get("/sign_in", user.SignIn)
	})
}
