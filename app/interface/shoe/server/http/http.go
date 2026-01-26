package http

import (
	"libong/common/server/http"
	"libong/login/auth"
	"shoe-manager/app/interface/shoe/service"
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
	group := server.Group("/manager/api/shoe")

	group.Use(auth.Authorize)
	group.POST("/add", addShoe)
	group.POST("/update", updateShoe)
	group.POST("/delete", deleteShoe)
	group.GET("/search/page", searchShoesPage)
	group.GET("/get", shoeById)
	group.POST("/update/hot", updateShoeHot)
	group.POST("/favour", changeShoeFavour)
	group.GET("/selectList/search", searchSelectList)
	return server
}
