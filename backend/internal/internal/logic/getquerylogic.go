// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"rank_list/internal/internal/svc"
	"rank_list/internal/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQueryLogic {
	return &GetQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQueryLogic) GetQuery(req *types.ListReq) (resp *types.ListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
