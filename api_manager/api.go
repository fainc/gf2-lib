package api_manager

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RebootReq struct {
	g.Meta   `path:"/manager/reboot" tags:"系统管理" method:"get" summary:"热重启系统"`
	FilePath string `json:"filePath" p:"filePath" v:"required"`
	Password string `json:"password" p:"password" v:"required"`
}

type RebootRes struct {
	Pid int `json:"pid"`
}
