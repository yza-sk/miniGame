package svc

import (
	"database/sql"
	"miniGame/internal/config"
	"miniGame/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	Config config.Config
	DB     *sql.DB
	Record *model.RecordModel
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	db, err := sql.Open("mysql", c.Database.DSN)
	if err != nil {
		return nil, err
	}
	recordModel := model.NewRecordModel(db)
	if err := recordModel.Init(); err != nil {
		return nil, err
	}
	return &ServiceContext{Config: c, DB: db, Record: recordModel}, nil
}
