// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"rank_list/internal/svc"
	"rank_list/internal/types"

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
	var respGradeList []types.Grade
	if req.Classify == "recent" {
		gradeList, err := l.svcCtx.Sql.FindRecent(l.ctx, req.Limit)
		if err != nil {
			return nil, err
		}

		for _, grade := range gradeList {
			respGrade := types.Grade{
				Id:       grade.Id,
				Name:     grade.Name,
				Score:    int(grade.Score),                             // 将int64转换为int
				Comment:  grade.Comment.String,                         // 将sql.NullString转换为string
				Finished: grade.Finished.Format("2006-01-02 15:04:05"), // 格式化时间
			}
			respGradeList = append(respGradeList, respGrade)
		}

	}else if req.Classify == "sort" {
		gradeList, err := l.svcCtx.Sql.FindTop(l.ctx, req.Limit)
		if err != nil {
			return nil, err
		}

		for _, grade := range gradeList {
			respGrade := types.Grade{
				Id:       grade.Id,
				Name:     grade.Name,
				Score:    int(grade.Score),                             // 将int64转换为int
				Comment:  grade.Comment.String,                         // 将sql.NullString转换为string
				Finished: grade.Finished.Format("2006-01-02 15:04:05"), // 格式化时间
			}
			respGradeList = append(respGradeList, respGrade)
		}
	}

	return &types.ListResp{GradeList: respGradeList}, nil
}
