package admin

import (
	"context"
	v1 "myblog/api/admin/v1"
)

var Account = &cAccount{}

type cAccount struct {
}

// info Get user infomation
func (c cAccount) Info(ctx context.Context, req *v1.AccountInfoReq) (res *v1.AccountInfoRes, err error) {
	return
}
