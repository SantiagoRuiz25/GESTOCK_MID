package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func ConectarPostgres() error {
	return orm.RegisterDataBase("default", "postgres", "connect_timeout=5")
}