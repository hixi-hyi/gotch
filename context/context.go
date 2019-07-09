package context

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

type key int

const (
	goonk key = 0
)

type Context struct{ context.Context }

func NewFromRequest(req *http.Request) Context {
	var ctx context.Context
	ctx = appengine.NewContext(req)
	ctx = context.WithValue(ctx, goonk, goon.FromContext(ctx))
	return Context{ctx}
}

func NewFromEchoContext(c echo.Context) Context {
	var ctx context.Context
	req := e.Request().(*standard.Request)
	ctx = appengine.WithContext(c.StdContext(), req.Request)
	ctx = context.WithValue(ctx, goonk, goon.FromContext(ctx))
	return Context{ctx}
}

func (ctx *Context) Goon() *goon.Goon {
	g, _ := ctx.Value(goonk).(*goon.Goon)
	return g
}

func (g *Gram) WithGoon(goon *goon.Goon) *Gram {
	ctx := context.WithValue(g, GOONK, goon)
	return &Gram{ctx}
}
