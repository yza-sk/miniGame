package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Database struct {
		// DSN is used for MySQL, e.g. user:pass@tcp(host:3306)/dbname?parseTime=true&loc=UTC
		DSN string
	}
	Cors struct {
		AllowedOrigins []string
	}
}
