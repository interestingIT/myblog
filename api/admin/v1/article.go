package v1

import (
	"myblog/internal/model"
	"myblog/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ArticleCreateReq struct {
	g.Meta `path:"/api/article/create" method:"post"`
	*model.ArticleInput
}

type ArticleCreateRes struct {
	LastId model.Id `json:"lastId"`
}

type ArticleDelReq struct {
	g.Meta `path:"/api/article/delete/{id}" method:"post"`
	*model.IdInput
	IsReal bool `dc:"Is permanently delete"`
}

type ArticleDelRes struct {
}

type ArticleUpdateReq struct {
	g.Meta `path:"/api/article/update/{id}" method:"post"`
	*model.IdInput
	*model.ArticleInput
}

type ArticleUpdateRes struct {
}

type ArticleListReq struct {
	g.Meta `path:"/api/article/list/*grpId" method:"get" sm:"search list"`
	*model.ArticleQuery
}

type ArticleListRes struct {
	List  []model.ArticleList `json:"list"`
	Total uint                `json:"total"`
}

type ArticleShowReq struct {
	g.Meta `path:"/api/article/show/{id}" method:"get"`
	*model.IdInput
}

type ArticleShowRes struct {
	*entity.Article
}
