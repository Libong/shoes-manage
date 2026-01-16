package conf

import (
	commonMysql "libong/common/orm/mysql"
	commonRedis "libong/common/redis"
	"libong/common/server/grpc"
	"libong/common/server/http"
	commonTool "libong/common/tool"
	loginConf "libong/login/app/interface/login/conf"
	"libong/login/app/interface/sms"
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
	Dao               *Dao
	RBACService       *grpc.Config
	SmsConfig         *sms.Config
	RoleServiceConfig *RoleServiceConfig
	OssService        *grpc.Config
	WechatConfig      *loginConf.WechatConfig
}
type RoleServiceConfig struct {
	CanOperateRoles []string
}
type Dao struct {
	Mysql *commonMysql.Config
	Redis *commonRedis.Config
}
type Server struct {
	HTTP *http.Config
	GRPC *grpc.Config
}
