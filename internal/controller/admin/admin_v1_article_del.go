package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"myblog/api/admin/v1"
)

func (c *ControllerV1) ArticleDel(ctx context.Context, req *v1.ArticleDelReq) (res *v1.ArticleDelRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
