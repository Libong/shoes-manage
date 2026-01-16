package http

import (
	"libong/common/server/http"
	loginConf "libong/login/app/interface/login/conf"
	loginHttp "libong/login/app/interface/login/server/http"
	loginService "libong/login/app/interface/login/service"
	"shoe-manager/app/interface/gateway/conf"

	accountConf "libong/rbac/app/interface/account/conf"
	accountHttp "libong/rbac/app/interface/account/server/http"
	accountService "libong/rbac/app/interface/account/service"

	roleConf "libong/rbac/app/interface/role/conf"
	roleHttp "libong/rbac/app/interface/role/server/http"
	roleService "libong/rbac/app/interface/role/service"
	shoeInterfaceConf "shoe-manager/app/interface/shoe/conf"
	shoeInterfaceHttp "shoe-manager/app/interface/shoe/server/http"
	shoeInterfaceService "shoe-manager/app/interface/shoe/service"

	categoryInterfaceConf "shoe-manager/app/interface/category/conf"
	categoryInterfaceHttp "shoe-manager/app/interface/category/server/http"
	categoryInterfaceService "shoe-manager/app/interface/category/service"

	dictInterfaceConf "shoe-manager/app/interface/dict/conf"
	dictInterfaceHttp "shoe-manager/app/interface/dict/server/http"
	dictInterfaceService "shoe-manager/app/interface/dict/service"
)

func New(conf *conf.Config) *http.Server {
	server := http.New(conf.Server.HTTP)

	dictInterfaceHttp.ConfigHttp(dictInterfaceService.New(&dictInterfaceConf.Service{}), server)

	categoryInterfaceHttp.ConfigHttp(categoryInterfaceService.New(&categoryInterfaceConf.Service{}), server)

	shoeInterfaceHttp.ConfigHttp(shoeInterfaceService.New(&shoeInterfaceConf.Service{
		OssService: conf.Service.OssService,
	}), server)
	loginHttp.ConfigHttp(loginService.New(&loginConf.Service{
		Dao: &loginConf.Dao{
			Redis: conf.Service.Dao.Redis,
			Mysql: conf.Service.Dao.Mysql,
		},
		RBACService:        conf.Service.RBACService,
		SmsConfig:          conf.Service.SmsConfig,
		LoginTokenExpireAt: conf.LoginTokenExpireAt,
		WechatConfig:       conf.Service.WechatConfig,
	}), server)
	accountHttp.ConfigHttp(accountService.New(&accountConf.Service{
		AccountService: conf.Service.RBACService,
	}), server)
	roleHttp.ConfigHttp(roleService.New(&roleConf.Service{
		AccountService: conf.Service.RBACService,
	}), server)
	return server
}
