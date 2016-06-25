package context

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"
)

type key int

const (
	goonk key = 0
)

type Gontext struct{ context.Context }

func NewFromRequest(req *http.Request) Gontext {
	var ctx context.Context
	ctx = appengine.NewContext(req)
	ctx = context.WithValue(ctx, goonk, goon.FromContext(ctx))
	return Gontext{ctx}
}

func NewFromEchoContext(c echo.Context) Gontext {
	var ctx context.Context
	req := c.Request().(*standard.Request)
	ctx = appengine.WithContext(c.Context(), req.Request)
	ctx = context.WithValue(ctx, goonk, goon.FromContext(ctx))
	return Gontext{ctx}
}

func (gtx *Gontext) Goon() *goon.Goon {
	g, _ := gtx.Value(goonk).(*goon.Goon)
	return g
}
