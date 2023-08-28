package v1

import (
	"myblog/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ArticleListReq struct {
	g.Meta `path:"/api/article/list/*grpId" method:"get"`
	*model.ArticleQuerySafe
}

type ArticleListRes struct {
	List  []model.ArticleListSafe `json:"list"`
	Total uint                    `json:"total"`
}

type ArticleRankReq struct {
	g.Meta `path:"/api/article/rank" method:"get" sm:"article rank"`
	Basis  int `v:"required|in:1,2" dc:"1-hot article 2-latest article"`
}

type ArticleRankRes struct {
	List []model.ArticleListSafe `json:"list"`
}

type ArticleShowReq struct {
	g.Meta `path:"/api/articl/show/{id}" method:"get"`
	*model.IdInput
}

type ArticleShowRes struct {
	*model.ArticleSafe
}

type ArticleHistReq struct {
	g.Meta `path:"/api/article/hist" method:"post" sm:"hist numbers"`
}

type ArticleHistRes struct {
	*model.IdInput
}
