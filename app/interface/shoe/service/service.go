package service

import (
	googleGrpc "google.golang.org/grpc"
	"libong/common/server/grpc"
	"shoe-manager/app/interface/shoe/conf"
	shoeServiceGrpc "shoe-manager/app/service/shoe/server/grpc"
	shoeServiceService "shoe-manager/app/service/shoe/service"
	ossServiceApi "shoe-manager/rpc/oss/api"
)

type Service struct {
	shoeService *shoeServiceGrpc.Server
	ossService  ossServiceApi.OssServiceClient
}

func New(c *conf.Service) *Service {
	var (
		ossConn *googleGrpc.ClientConn
		err     error
	)

	service := &Service{
		shoeService: shoeServiceGrpc.New(shoeServiceService.New()),
	}

	ossConn, err = grpc.NewConnection(c.OssService)
	if err != nil {
		panic(err)
	}
	service.ossService = ossServiceApi.NewOssServiceClient(ossConn)
	return service
}
