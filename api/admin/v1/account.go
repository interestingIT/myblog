package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"login" method:"post" sm:"login" tags:"account"`
	Username string `v:"required|length: 3,20"`
	Password string `v:"required|length: 8,30"`
}

type LoginRes struct {
}
