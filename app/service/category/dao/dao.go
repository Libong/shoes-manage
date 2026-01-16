package dao

import (
	"libong/common/orm/mysql"
)

// Dao .
type Dao struct {
	*commonMysql.Client
}

// New .
func New() *Dao {
	if commonMysql.MysqlClient == nil {
		panic("MysqlClient is nil")
	}
	return &Dao{
		Client: commonMysql.MysqlClient,
	}
}
