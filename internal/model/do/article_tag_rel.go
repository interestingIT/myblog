// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleTagRel is the golang structure of table article_tag_rel for DAO operations like Where/Data.
type ArticleTagRel struct {
	g.Meta    `orm:"table:article_tag_rel, do:true"`
	Id        interface{} // id
	ArticleId interface{} // 文章id
	TagId     interface{} // 标签id
}
