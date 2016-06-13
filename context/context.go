package context

import (
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

func (gtx *Gontext) Goon() *goon.Goon {
	g, _ := gtx.Value(goonk).(*goon.Goon)
	return g
}
