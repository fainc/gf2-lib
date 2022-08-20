package response

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"
)

// HandlerResponse 默认数据返回中间件
func HandlerResponse(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	// api.json 不作处理
	if r.RequestURI == "/api.json" {
		return
	}
	// 已有自定义输出内容，不作处理
	if r.Response.BufferLength() > 0 && gmeta.Get(res, "mime").String() == "custom" {
		return
	}
	if err != nil {
		if code.Code() == 50 || code.Code() == 52 { // 服务器错误
			Json().ServerError(ctx)
			return
		}
		if code.Code() == 401 { // 登录
			Json().Authorization(ctx, code.Message(), code.Detail())
			return
		}
		Json().Error(ctx, err.Error(), code.Code(), code.Detail()) // 常规错误
		return
	}
	if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		switch r.Response.Status {
		case http.StatusNotFound:
			Json().NotFound(ctx)
			return
		case http.StatusUnauthorized:
			return
		default:
			Json().ServerError(ctx)
			return
		}
	}
	Json().Success(ctx, res)
}
