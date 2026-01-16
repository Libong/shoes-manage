package service

import (
	"shoe-manager/app/interface/category/conf"
	categoryServiceGrpc "shoe-manager/app/service/category/server/grpc"
	categoryServiceService "shoe-manager/app/service/category/service"
)

type Service struct {
	categoryService *categoryServiceGrpc.Server
}

func New(c *conf.Service) *Service {
	service := &Service{
		categoryService: categoryServiceGrpc.New(categoryServiceService.New()),
	}
	return service
}
