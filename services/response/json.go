package response

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

func JSON(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetContentType("application/json; charset=utf-8")

	jbyte, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(ctx, string(jbyte))

}
