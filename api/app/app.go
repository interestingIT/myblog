// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package app

import (
	"context"
	
	"myblog/api/app/v1"
)

type IAppV1 interface {
	ArticleList(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error)
	ArticleRank(ctx context.Context, req *v1.ArticleRankReq) (res *v1.ArticleRankRes, err error)
	ArticleShow(ctx context.Context, req *v1.ArticleShowReq) (res *v1.ArticleShowRes, err error)
	ArticleHist(ctx context.Context, req *v1.ArticleHistReq) (res *v1.ArticleHistRes, err error)
}


