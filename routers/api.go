package routers

import (
	ctrlIndex "api-teste/controllers/index"
	ctrlUser "api-teste/controllers/user"

	"github.com/roger-russel/fasthttp-router-middleware/pkg/middleware"

	routers "github.com/fasthttp/router"
	api "github.com/valyala/fasthttp" //servidor http - bem mais rápida que a nativa
)

//Get Retorno da rota
func Get() func(ctx *api.RequestCtx) {
	m := middleware.New([]middleware.Middleware{exampleAuth})

	r := routers.New()
	r.GET("/", m(ctrlIndex.Index))
	r.GET("/user", ctrlUser.Index)
	return r.Handler
}

func exampleAuth(ctx *api.RequestCtx) bool {
	//ctx.Response.SetStatusCode(api.StatusUnauthorized) //troco status
	return true
}
