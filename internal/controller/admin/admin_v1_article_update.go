package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"myblog/api/admin/v1"
)

func (c *ControllerV1) ArticleUpdate(ctx context.Context, req *v1.ArticleUpdateReq) (res *v1.ArticleUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
