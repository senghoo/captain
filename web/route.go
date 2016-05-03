package web

import (
	"github.com/go-macaron/binding"
	"github.com/senghoo/captain/web/controllers"
	"github.com/senghoo/captain/web/controllers/user"
	"github.com/senghoo/captain/web/middleware"
)

func (s *Server) initRoute() {
	reqSignIn := middleware.Toggle(&middleware.ToggleOptions{SignInRequire: true})
	// ignSignIn := middleware.Toggle(&middleware.ToggleOptions{SignInRequire: setting.Service.RequireSignInView})
	// ignSignInAndCsrf := middleware.Toggle(&middleware.ToggleOptions{DisableCsrf: true})
	reqSignOut := middleware.Toggle(&middleware.ToggleOptions{SignOutRequire: true})

	s.m.Get("/", reqSignIn, controllers.Main)
	s.m.Group("/user", func() {
		s.m.Get("/sign_in", reqSignOut, user.SignIn)
		s.m.Post("/sign_in", reqSignOut, binding.BindIgnErr(user.SignInForm{}), user.SignInPost)
	})
}
