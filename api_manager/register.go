package api_manager

import (
	"context"
	"errors"
	"os"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// RemoteRegister 部署者远程注册
func RemoteRegister() (err error) {
	host, _ := os.Hostname()
	deployerAppId, err := g.Cfg().Get(context.Background(), "server.deployerAppId")
	if err != nil {
		return
	}
	if deployerAppId.String() == "" {
		err = errors.New("未设置deployerAppId")
		return
	}
	deployerAppSecret, err := g.Cfg().Get(context.Background(), "server.deployerAppSecret")
	if err != nil {
		return
	}
	if deployerAppSecret.String() == "" {
		err = errors.New("未设置deployerAppSecret")
		return
	}
	pid := os.Getpid()
	path, _ := os.Executable()
	signStr := "appPath=" + path + "&deployerAppId=" + deployerAppId.String() + "&host=" + host + "&pid=" + gconv.String(pid) + "&faDeployerKey=" + deployerAppSecret.String()
	signMd5, _ := gmd5.EncryptString(signStr)
	reg := g.Map{
		"appPath":       path,
		"host":          host,
		"deployerAppId": deployerAppId.String(),
		"pid":           pid,
		"sign":          signMd5,
	}
	url := "https://api.deployer.fain.cn/open-api/register/go-app"
	_ = g.Client().PostContent(context.Background(), url, reg)
	return
}
