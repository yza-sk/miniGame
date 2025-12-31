// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"database/sql"

	"rank_list/internal/model"
	"rank_list/internal/svc"
	"rank_list/internal/types"

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
	grade := &model.Grade{
		Name:    req.Name,
		Score:   int64(req.Score),
		Comment: sql.NullString{String: req.Comment, Valid: req.Comment != ""},
	}
	_, err = l.svcCtx.Sql.Insert(l.ctx, grade)

	if err != nil {
		return nil, err
	}
	return
}
