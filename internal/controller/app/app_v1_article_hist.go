package app

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"myblog/api/app/v1"
)

func (c *ControllerV1) ArticleHist(ctx context.Context, req *v1.ArticleHistReq) (res *v1.ArticleHistRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
