package http

import (
	"net/http"

	"github.com/brettmostert/fnple-go/go/components/ledger/internal/common"
)

type api struct {
	ctx    *common.AppContext
	router *http.ServeMux
}

func NewApi(ctx *common.AppContext) *api {
	s := &api{
		ctx:    ctx,
		router: http.NewServeMux(),
	}
	s.routes()
	return s
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
