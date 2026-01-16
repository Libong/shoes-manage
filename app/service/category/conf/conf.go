package conf

import (
	"libong/common/orm/mysql"
	"libong/common/server/grpc"
	commonTool "libong/common/tool"
)

type Config struct {
	Server  *Server
	Service *Service
}

func New() *Config {
	conf := &Config{}
	//初始化配置文件
	commonTool.GrabConfigFile(conf)
	return conf
}

type Service struct {
	Dao *Dao
}
type Dao struct {
	Mysql *commonMysql.Config
}
type Server struct {
	//HTTP *http
	GRPC *grpc.Config
}
