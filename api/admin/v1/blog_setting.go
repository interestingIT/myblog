/*
* author: faith
* date: 2023-10-9
* Description: blog setting
 */
package v1

import (
	"myblog/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminBlogSettingUpdateReq struct {
	g.Meta `path:"/api/admin/blog/setting/update" method:"post"`
	*entity.BlogSetting
}

type AdminBlogSettingUpdateRes struct {
}

type AdminBlogSettingDetailReq struct {
	g.Meta `path:"/api/admin/blog/setting/detail" method:"get"`
}

type AdminBlogSettingDetailRes struct {
	*entity.BlogSetting
}
