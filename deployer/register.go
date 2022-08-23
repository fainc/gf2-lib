package deployer

import (
	"context"
	"errors"
	"os"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// RemoteRegister 部署中心远程注册
// deployerCenterHost 部署中心远端域名定义 一般为 https://api.deployer.fain.cn
// deployerConnectHost 本机通讯域名/IP 用于部署中心与本机通讯（心跳检查、更新通知、热更新等），不要使用集群或负载均衡域名，如应用多机部署请使用每个IP+端口号形式注册
// deployerAppId 部署中心分配 deployerAppId
// deployerAppSecret 部署中心分配 deployerAppSecret

func RemoteRegister() (err error) {
	host, _ := os.Hostname()
	deployerHeartbeatUrl, err := g.Cfg().Get(context.Background(), "server.deployerConnectHost")
	if err != nil {
		return
	}
	if deployerHeartbeatUrl.String() == "" {
		err = errors.New("未设置deployerConnectHost")
		return
	}
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
