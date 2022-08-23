package deployer

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RebootReq struct {
	g.Meta   `path:"/deployer-api/reboot" tags:"部署中心" method:"get" summary:"热重启系统"`
	FilePath string `json:"filePath" p:"filePath" v:"required#参数错误"`
	Sign     string `json:"sign" p:"sign" v:"required#参数错误"`
}
type RebootRes struct {
	Pid int `json:"pid"`
}
