// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ArticleContent is the golang structure for table article_content.
type ArticleContent struct {
	Id        uint64 `json:"id"        description:"文章内容id"`
	ArticleId int64  `json:"articleId" description:"文章id"`
	Content   string `json:"content"   description:"正文内容"`
}
