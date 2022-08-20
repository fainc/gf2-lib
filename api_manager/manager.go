package api_manager

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/fainc/gf2-lib/response"
)

var Manager = cManager{}

type cManager struct{}

func (c *cManager) Reboot(ctx context.Context, req *RebootReq) (res *RebootRes, err error) {
	res = &RebootRes{
		Pid: os.Getpid(),
	}
	password, err := g.Cfg().Get(ctx, "server.gracefulPassword")
	if err != nil {
		return
	}
	if password.String() == "" {
		err = response.StandardError(ctx, -101, "未配置服务端密码", nil)
		return
	}
	serverPassword, err := gmd5.EncryptString(password.String())
	if err != nil {
		return
	}
	userPassword, err := gmd5.EncryptString(req.Sign)
	if err != nil {
		return
	}
	if serverPassword != userPassword {
		err = response.StandardError(ctx, -102, "密码无效", nil)
		return
	}
	err = ghttp.RestartAllServer(ctx, req.FilePath)
	return
}
