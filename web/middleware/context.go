package middleware

import "gopkg.in/macaron.v1"

type Context struct {
}

func Contexter() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &Context{}
		c.Map(ctx)
	}
}
