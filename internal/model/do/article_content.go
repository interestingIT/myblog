// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleContent is the golang structure of table article_content for DAO operations like Where/Data.
type ArticleContent struct {
	g.Meta    `orm:"table:article_content, do:true"`
	Id        interface{} // 文章内容id
	ArticleId interface{} // 文章id
	Content   interface{} // 正文内容
}
