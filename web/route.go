package web

import (
	"github.com/go-macaron/binding"
	"github.com/senghoo/captain/web/controllers"
	"github.com/senghoo/captain/web/controllers/docker"
	"github.com/senghoo/captain/web/controllers/github"
	"github.com/senghoo/captain/web/controllers/repo"
	"github.com/senghoo/captain/web/controllers/user"
	"github.com/senghoo/captain/web/controllers/workspace"
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
		s.m.Get("/", github.Info)
		s.m.Get("/auth", github.Auth)
		s.m.Get("/callback", github.Callback)
		s.m.Get("/webhook/:id([0-9]+)", github.Webhook)
	}, reqSignIn)

	s.m.Group("/docker", func() {
		s.m.Get("/", docker.Index)
		s.m.Get("/new", docker.New)
		s.m.Post("/new", binding.BindIgnErr(docker.NewForm{}), docker.NewPost)
		s.m.Get("/:id([0-9]+)", docker.Info)
	}, reqSignIn)

	s.m.Group("/workspace", func() {
		s.m.Get("/", workspace.Index)
		s.m.Get("/new", workspace.New)
		s.m.Post("/new", binding.BindIgnErr(workspace.NewForm{}), workspace.NewPost)
		s.m.Get("/:id([0-9]+)", workspace.Info)
		s.m.Get("/:id([0-9]+)/repository/new", workspace.AddRepository)
		s.m.Post("/:id([0-9]+)/repository/new", binding.BindIgnErr(workspace.AddRepositoryForm{}), workspace.PostAddRepository)
		s.m.Get("/:id([0-9]+)/workflow/new", workspace.AddWorkflow)
		s.m.Post("/:id([0-9]+)/workflow/new", binding.BindIgnErr(workspace.AddWorkflowForm{}), workspace.PostAddWorkflow)
	}, reqSignIn)

	s.m.Group("/repo", func() {
		s.m.Get("/:id([0-9]+)/build", repo.Build)
	})

}
