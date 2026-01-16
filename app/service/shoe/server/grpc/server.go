package grpc

import (
	"libong/common/server/grpc"
	"shoe-manager/app/service/shoe/api"
	"shoe-manager/app/service/shoe/service"
)

type Server struct {
	service *service.Service
	api.UnimplementedShoeServiceServer
}

// New 用于引用rpc service
func New(svr *service.Service) *Server {
	return &Server{
		service: svr,
	}
}

// NewServer 用于单独启动rpc service
func NewServer(svc *service.Service, conf *grpc.Config) *grpc.Server {
	//设置ip和端口号
	s := grpc.New(conf)
	server := &Server{service: svc}
	//将服务注册为rpc服务
	api.RegisterShoeServiceServer(s.Server(), server)
	//心跳注册
	//s.RegisterHealth(server)
	return s
}
