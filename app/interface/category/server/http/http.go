package http

import (
	"libong/common/server/http"
	"libong/login/auth"
	"shoe-manager/app/interface/category/service"
)

// http层的service对象
var svc *service.Service

// NewServer 初始化
func NewServer(s *service.Service, c *http.Config) *http.Server {
	//初始化http服务对象
	server := http.New(c)
	//http配置路径、中间件
	ConfigHttp(s, server)
	return server
}

func ConfigHttp(s *service.Service, server *http.Server) *http.Server {
	//提出service对象 用于controller调用
	svc = s
	//路径配置
	group := server.Group("/manager/api/category")

	group.Use(auth.Authorize)
	group.POST("/add", addCategory)
	group.POST("/update", updateCategory)
	group.POST("/delete", deleteCategory)
	group.GET("/search/page", searchCategoriesPage)
	group.GET("/get", categoryById)
	return server
}
