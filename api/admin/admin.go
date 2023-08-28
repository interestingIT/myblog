// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"context"
	
	"myblog/api/admin/v1"
)

type IAdminV1 interface {
	ArticleCreate(ctx context.Context, req *v1.ArticleCreateReq) (res *v1.ArticleCreateRes, err error)
	ArticleDel(ctx context.Context, req *v1.ArticleDelReq) (res *v1.ArticleDelRes, err error)
	ArticleUpdate(ctx context.Context, req *v1.ArticleUpdateReq) (res *v1.ArticleUpdateRes, err error)
	ArticleList(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error)
	ArticleShow(ctx context.Context, req *v1.ArticleShowReq) (res *v1.ArticleShowRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}


