// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	// "database/sql"
	"fmt"
	"rank_list/internal/config"
	"rank_list/internal/entity"
	"rank_list/internal/model"

	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	MysqlConn sqlx.SqlConn
	Sql    model.GradeModel
}


func NewServiceContext(c config.Config) *ServiceContext {
	client, err := NewNacosClient()
	if err != nil {
		fmt.Printf("NewNacosClient err: %v", err)
	}
	var GetConfigParam = entity.ConfigParam{
		Group: "MINIGAME",
		DataId: "mysql",
		ConfigType: config.MysqlConfig{},
	}
    mysqlConfigInter, err := client.GetConfig(GetConfigParam)

	var DBconf config.DB
	mapstructure.Decode(mysqlConfigInter, &DBconf)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// fmt.Printf("%s", DBconf.Dsn)

	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// dbconn, err := sql.Open(DBconf.DriverName, DBconf.Dsn)
	// if err != nil {
	// 	fmt.Printf("sql.Open err: %v", err)
	// }

	// mysqlconn := sqlx.NewSqlConnFromDB(dbconn)
	mysqlconn := sqlx.NewSqlConn(DBconf.DriverName, DBconf.Dsn)

	mysql := model.NewGradeModel(mysqlconn)

	return &ServiceContext{
		Config: c,
		MysqlConn: mysqlconn,
		Sql:    mysql,
	}
}
