package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/api/user/login" method:"post" tags:"Account"`
	Username string `v:"required|length: 3,20"`
	Password string `v:"required|length: 8,30"`
}

type LoginRes struct {
	Token string
}

type LogoutReq struct {
	g.Meta `path:"/api/user/logout" method:"post"`
	Token  string
}

type LogoutRes struct {
	Result bool
}
