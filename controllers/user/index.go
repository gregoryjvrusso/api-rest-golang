package user

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Index router controller
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "User Index")
}
