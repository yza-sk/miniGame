// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"rank_list/internal/internal/svc"
	"rank_list/internal/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitLogic {
	return &SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req *types.SubmitReq) (resp *types.SubmitResp, err error) {
	// todo: add your logic here and delete this line

	return
}
