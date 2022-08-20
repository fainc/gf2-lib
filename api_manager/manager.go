package api_manager

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

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
	signStr := "filePath=" + req.FilePath + "&pid=" + gconv.String(res.Pid) + "&password=" + password.String()
	serverSign, err := gmd5.EncryptString(signStr)
	if err != nil {
		return
	}
	if serverSign != req.Sign {
		err = response.StandardError(ctx, -102, "签名无效", nil)
		return
	}
	err = ghttp.RestartAllServer(ctx, req.FilePath)
	return
}
