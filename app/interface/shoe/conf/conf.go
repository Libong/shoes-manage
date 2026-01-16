package conf

import (
	commonMysql "libong/common/orm/mysql"
	commonRedis "libong/common/redis"
	"libong/common/server/grpc"
	"libong/common/server/http"
	commonTool "libong/common/tool"
)

type Config struct {
	Server             *Server
	Service            *Service
	LoginTokenExpireAt int64
}

func New() *Config {
	conf := &Config{}
	//初始化配置文件
	commonTool.GrabConfigFile(conf)
	return conf
}

type Service struct {
	OssService *grpc.Config
}
type Dao struct {
	Mysql *commonMysql.Config
	Redis *commonRedis.Config
}
type Server struct {
	HTTP *http.Config
	GRPC *grpc.Config
}
