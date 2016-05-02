package web

import (
	"github.com/go-macaron/binding"
	"github.com/senghoo/captain/web/controllers"
	"github.com/senghoo/captain/web/controllers/user"
)

func (s *Server) initRoute() {
	s.m.Get("/", controllers.Main)
	s.m.Group("/user", func() {
		s.m.Get("/sign_in", user.SignIn)
		s.m.Post("/sign_in", binding.Bind(user.SignInForm{}), user.SignInPost)
	})
}
