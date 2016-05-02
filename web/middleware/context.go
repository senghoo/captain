package middleware

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/senghoo/captain/models"
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

func Contexter() macaron.Handler {
	return func(c *macaron.Context, cache cache.Cache, sess session.Store, f *session.Flash, x csrf.CSRF) {
		ctx := &Context{
			Cache:   cache,
			csrf:    x,
			Flash:   f,
			Session: sess,
			User:    models.GetUserFromSession(sess),
		}
		c.Map(ctx)
	}
}
