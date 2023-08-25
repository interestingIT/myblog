package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginReq struct {
	g.Meta   `path:"login" method:"post" sm:"login" tags:"account"`
	Username string `v:"required|length: 3,20"`
	Password string `v:"required|length: 8,30"`
}

type LoginRes struct {
	Token string `sm:"verify" dc:"In API, Authorization:token"`
}

type LogoutReq struct {
	g.Meta `path:"logout" method:"post" sm:"logout" tags:"account"`
	Token  string
}

type LogoutRes struct {
	Result bool
}

type AccountInfoReq struct {
	g.Meta `path:"info" method:"get" sm:"get login info" tags:"Account"`
}

type AccountInfoRes struct {
	Username  string      `json:"username" dc:"user name"`
	Nickname  string      `json:"nickname" dc:"nick name"`
	Avator    string      `json:"avator" dc:"avator"`
	Register  string      `json:"register" dc:"register time"`
	LastLogin *gtime.Time `json:"lastLogin" dc:"last login time"`
}
