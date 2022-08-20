package response

import (
	"context"
	"os"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

type json struct{}

var jsonVar = json{}

func Json() *json {
	return &jsonVar
}

func (rec *json) Success(ctx context.Context, data interface{}) {
	rec.Writer(ctx, data, "success", 200, 200, 0, nil)
}

func (rec *json) Error(ctx context.Context, message string, errCode int, ext interface{}) {
	rec.Writer(ctx, nil, message, 400, 400, errCode, ext)
}

func (rec *json) ServerError(ctx context.Context) {
	rec.Writer(ctx, nil, "内部服务器错误", 500, 500, 500, nil)
}

func (rec *json) Authorization(ctx context.Context, message string, ext interface{}) {
	rec.Writer(ctx, nil, message, 401, 401, 401, ext)
}

func (rec *json) NotFound(ctx context.Context) {
	rec.Writer(ctx, nil, "Not Found", 404, 404, 404, nil)
}

type JsonFormat struct {
	Code      int         `json:"code"`      // 业务码，用于业务判断（首选），固定 200/400/401/404/500，与http状态码同步，通常取该值判断是否有错误需要处理
	ErrorCode int         `json:"errorCode"` // 错误码，用于业务判断（可选），-1/400(通用错误)/51(参数验证错误)/401/404/500/other，通常忽略该值，除非业务需要判断详细错误类型（例：交易场景，交易失败返回400业务码时，返回余额不足、账户冻结等详细错误码用于后续业务处理）
	Message   interface{} `json:"message"`   // 消息，业务码非400时固定输出 Success/Authorization/NotFound/InternalError中的一种，400时输出详细错误（可能含i18n转译，建议仅展示或记录信息，不可用于业务判断）
	Data      interface{} `json:"data"`      // 返回数据
	TraceId   interface{} `json:"traceId"`   // 请求唯一追踪ID,用于日志定位
	Server    string      `json:"server"`    // 服务端 host name md5 值
	Runtime   int64       `json:"runtime"`   // 当前任务执行耗时（ms）
	Ext       interface{} `json:"ext"`       // 拓展数据（可能含有多个错误详情或其他附加数据，例：复杂登录场景下的401错误返回具体登录地址）
}

// Writer 数据输出
func (rec *json) Writer(ctx context.Context, data interface{}, message string, status int, code int, errCode int, ext interface{}) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteStatus(status)
	r.Response.ClearBuffer()
	serverName, _ := os.Hostname()
	serverName, _ = gmd5.Encrypt(serverName)
	r.Response.WriteJson(JsonFormat{
		Code:      code,
		Message:   message,
		Data:      data,
		TraceId:   gctx.CtxId(ctx),
		Server:    strings.ToUpper(serverName),
		Ext:       ext,
		ErrorCode: errCode,
		Runtime:   gtime.Now().TimestampMilli() - r.EnterTime,
	})
	r.Response.Header().Set("Content-Type", "application/json;charset=utf-8")
	if status == 401 {
		r.ExitAll()
	}
}
