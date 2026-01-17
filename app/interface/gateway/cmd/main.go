package main

import (
	"libong/common/log"
	commonMysql "libong/common/orm/mysql"
	commonRedis "libong/common/redis"
	"libong/common/server"
	"libong/common/snowflake"
	commonToolCrypt "libong/common/tool/crypt"
	"shoe-manager/app/interface/gateway/conf"
	"shoe-manager/app/interface/gateway/server/http"
)

func main() {
	//初始化配置文件
	config := conf.New()
	commonToolCrypt.LoadPrivateKey()
	log.Init()
	//初始化雪花算法 用于生成id
	snowflake.NewWorker(snowflake.WorkerIDBits, snowflake.DataCenterIDBits)
	//初始化数据库
	commonMysql.NewClient(config.Service.Dao.Mysql)
	commonRedis.NewClient(config.Service.Dao.Redis)
	initGlobal(config)
	//启动服务
	httpServer := http.New(config)
	server.StartWaitingForQuit(httpServer)
}
func initGlobal(conf *conf.Config) {
}
