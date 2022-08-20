package response

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// CustomRes 自定义返回数据标注，使用该类型数据返回时，全局后置中间件ResponseHandler将不再处理返回数据，请自行提前输出
type CustomRes struct {
	g.Meta `mime:"custom" sm:"自定义数据返回" dc:"本接口使用自定义数据返回，非OPEN API v3规范，具体返回数据字段请联系管理员获取"`
	Data   interface{} `json:"data"`
}

// AuthorizedError 授权
func AuthorizedError(ctx context.Context, message string, ext interface{}) error {
	// ctx 预留上下文翻译
	return gerror.NewCode(gcode.New(401, message, ext))
}

// StandardError 标准错误
func StandardError(ctx context.Context, code int, message string, ext interface{}) error {
	if code == 401 {
		return AuthorizedError(ctx, message, ext)
	}
	return gerror.NewCode(gcode.New(code, message, ext))
}
