// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DriverConf MysqlConfig
	DB DB `json:"db"`
}

type MysqlConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Username string `json:"user"`
	Password string `json:"pass"`
	Database string `json:"databases"`
	Dsn string `json:"dsn"`
}

type DB struct {
	Dsn string `json:"dsn"`
	DriverName string `json:"driver_name"`
}