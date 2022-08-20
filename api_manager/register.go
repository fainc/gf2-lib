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
	faDeployerId, err := g.Cfg().Get(context.Background(), "server.faDeployerId")
	if err != nil {
		return
	}
	if faDeployerId.String() == "" {
		err = errors.New("未设置faDeployerId")
		return
	}
	faDeployerKey, err := g.Cfg().Get(context.Background(), "server.faDeployerKey")
	if err != nil {
		return
	}
	if faDeployerKey.String() == "" {
		err = errors.New("未设置faDeployerKey")
		return
	}
	pid := os.Getpid()
	signStr := "faDeployerId=" + faDeployerId.String() + "&host=" + host + "&pid=" + gconv.String(pid) + "&faDeployerKey=" + faDeployerKey.String()
	signMd5, _ := gmd5.EncryptString(signStr)
	reg := g.Map{
		"host":         host,
		"faDeployerId": faDeployerId.String(),
		"pid":          pid,
		"sign":         signMd5,
	}
	g.Dump(reg)
	return
}
