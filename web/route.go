package web

import (
	"github.com/go-macaron/binding"
	"github.com/senghoo/captain/web/controllers"
	"github.com/senghoo/captain/web/controllers/docker"
	"github.com/senghoo/captain/web/controllers/github"
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
		s.m.Get("/sign_out", user.SignOut)
	})
	s.m.Group("/github", func() {
		s.m.Get("/", github.List)
		s.m.Get("/auth", github.Auth)
		s.m.Get("/callback", github.Callback)
	}, reqSignIn)

	s.m.Group("/docker", func() {
		s.m.Get("/", docker.Index)
	}, reqSignIn)
}
