package service

import (
	"shoe-manager/app/interface/dict/conf"
	dictServiceGrpc "shoe-manager/app/service/dict/server/grpc"
	dictServiceService "shoe-manager/app/service/dict/service"
)

type Service struct {
	dictService *dictServiceGrpc.Server
}

func New(c *conf.Service) *Service {
	service := &Service{
		dictService: dictServiceGrpc.New(dictServiceService.New()),
	}
	return service
}
