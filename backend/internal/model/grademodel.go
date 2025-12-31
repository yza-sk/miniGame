package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GradeModel = (*customGradeModel)(nil)

type (
	// GradeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGradeModel.
	GradeModel interface {
		gradeModel
		withSession(session sqlx.Session) GradeModel
		FindRecent(ctx context.Context, limit int) ([]*Grade, error)
		FindTop(ctx context.Context, limit int) ([]*Grade, error)
	}

	customGradeModel struct {
		*defaultGradeModel
	}
)

// NewGradeModel returns a model for the database table.
func NewGradeModel(conn sqlx.SqlConn) GradeModel {
	return &customGradeModel{
		defaultGradeModel: newGradeModel(conn),
	}
}

func (m *customGradeModel) withSession(session sqlx.Session) GradeModel {
	return NewGradeModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customGradeModel) FindRecent(ctx context.Context, limit int) ([]*Grade, error) {
	var grades []*Grade
	// 假设有自增id或时间戳字段，按此字段降序排列获取最后100条
	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d", gradeRows, m.table, limit)
	err := m.conn.QueryRowsCtx(ctx, &grades, query)
	if err != nil {
		return nil, err
	}
	return grades, nil
}

func (m *customGradeModel) FindTop(ctx context.Context, limit int) ([]*Grade, error) {
	var grades []*Grade
	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY score DESC LIMIT %d", gradeRows, m.table, limit)
	err := m.conn.QueryRowsCtx(ctx, &grades, query)
	if err != nil {
		return nil, err
	}
	return grades, nil
}
