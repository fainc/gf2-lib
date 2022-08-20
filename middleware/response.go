package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/fainc/gf2-lib/response"
)

func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	response.HandlerResponse(r)
}
