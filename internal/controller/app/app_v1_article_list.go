package app

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"myblog/api/app/v1"
)

func (c *ControllerV1) ArticleList(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
