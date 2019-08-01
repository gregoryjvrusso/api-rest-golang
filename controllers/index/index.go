package index

import (
	"api-teste/services/response"

	"github.com/valyala/fasthttp"
)

type Data struct {
	Foo string `json:"foo"`
	Bar string `json:"bar,omitempty"`
}

// Index router controller
func Index(ctx *fasthttp.RequestCtx) {
	d := Data{
		Foo: "foo",
		Bar: "",
	}

	response.JSON(ctx, &d)

}
