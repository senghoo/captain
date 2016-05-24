package middleware

import (
	"fmt"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/modules/settings"
	"gopkg.in/macaron.v1"
)

type Context struct {
	*macaron.Context
	Cache   cache.Cache
	csrf    csrf.CSRF
	Flash   *session.Flash
	Session session.Store

	User     *models.User
	IsSigned bool
}

func (c *Context) SetUser(u *models.User) {
	c.Session.Set("uid", u.ID)
}

func (c *Context) GetUser() *models.User {
	uid := c.Session.Get("uid")
	if uid == nil {
		return nil
	}

	if id, ok := uid.(int64); ok {
		return models.GetUserByID(id)
	}
	return nil
}

func (c *Context) Logout() {
	c.User = nil
	c.IsSigned = false
	c.Session.Delete("uid")
}

// HasError returns true if error occurs in form validation.
func (ctx *Context) HasError() bool {
	hasErr, ok := ctx.Data["HasError"]
	if !ok {
		return false
	}
	ctx.Flash.ErrorMsg = ctx.Data["ErrorMsg"].(string)
	ctx.Data["Flash"] = ctx.Flash
	return hasErr.(bool)
}

func (ctx *Context) HTML(status int, name string, data ...interface{}) {
	ctx.Context.HTML(status, name, data...)
}

func (ctx *Context) HandleErr(err error, ret string) {
	fmt.Printf("500: %s\n", err)
	if err != nil && macaron.Env != macaron.PROD {
		ctx.Data["ErrorMsg"] = err
	}
	ctx.Data["Ret"] = ret
	ctx.HTML(500, "500")
}

func (ctx *Context) NotFound(msg string) {
	ctx.HTML(404, "404")
}

func Contexter() macaron.Handler {
	return func(c *macaron.Context, cache cache.Cache, sess session.Store, f *session.Flash, x csrf.CSRF) {
		ctx := &Context{
			Context: c,
			Cache:   cache,
			csrf:    x,
			Flash:   f,
			Session: sess,
		}
		user := ctx.GetUser()
		if user != nil {
			ctx.User = user
			ctx.IsSigned = true
		}
		c.Map(ctx)
	}
}

type ToggleOptions struct {
	SignInRequire  bool
	SignOutRequire bool
	AdminRequire   bool
	DisableCsrf    bool
}

func Toggle(options *ToggleOptions) macaron.Handler {
	return func(ctx *Context) {

		// Redirect to dashboard if user tries to visit any non-login page.
		if options.SignOutRequire && ctx.IsSigned && ctx.Req.RequestURI != "/" {
			ctx.Redirect(settings.GetOrDefault("site.url", "") + "/")
			return
		}

		if !options.SignOutRequire && !options.DisableCsrf && ctx.Req.Method == "POST" {
			csrf.Validate(ctx.Context, ctx.csrf)
			if ctx.Written() {
				return
			}
		}

		if options.SignInRequire && !ctx.IsSigned {
			ctx.Redirect(settings.GetOrDefault("site.url", "") + "/user/sign_in")
			return
		}
		if options.SignOutRequire && ctx.IsSigned {
			ctx.Redirect(settings.GetOrDefault("site.url", "") + "/")
			return
		}

	}
}
