package model

import "myblog/internal/model/entity"

type ArticleInput struct {
	GroupId uint32 `json:"groupId" v:"required|integer|between:1,4294967295"`
	Title   string `json:"title" v:"required|length:2, 100"`
	Autor   string `json:"author" v:"length:2,30"`
	Tags    string `json:"tags" v:"length:2, 200"`
	Content string `json:"content" v:"length:2, 1000000"`
}

type ArticleQuery struct {
	ArticleQuerySafe
	Onshow bool `json:"onshow" dc:"是否查询只发布的文章"`
	IsDel  bool `json:"isDel" dc:"是否查询删除掉的文章"`
}

type ArticleQuerySafe struct {
	Paging
	GrpId  Id     `v:"integer|between:1,4294967295" json:"grpId"`
	Search string `v:"length: 1,30" json:"search" dc:"查询文本，会检索标题、标签、简介"`
}

type ArticleList struct {
	entity.Article
	Content   Out `json:"content,omitempty"`
	DeletedAt Out `json:"deletedAt,omitempty"`
}

type ArticleListSafe struct {
	entity.Article
	ArticleSafe
	DeletedAt Out `json:"deletedAt,omitempty"`
}

type ArticleSafe struct {
	*entity.Article
	Order     Out `json:"order,omitempty"`
	Ontop     Out `json:"ontop,omitempty"`
	Onshow    Out `json:"onshow,omitempty"`
	DeletedAt Out `json:"deletedAt,omitempty"`
}
