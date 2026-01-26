package service

import (
	googleGrpc "google.golang.org/grpc"
	"libong/common/server/grpc"
	"shoe-manager/app/interface/shoe/conf"
	dictServiceGrpc "shoe-manager/app/service/dict/server/grpc"
	dictServiceService "shoe-manager/app/service/dict/service"
	shoeServiceGrpc "shoe-manager/app/service/shoe/server/grpc"
	shoeServiceService "shoe-manager/app/service/shoe/service"
	ossServiceApi "shoe-manager/rpc/oss/api"
)

type Service struct {
	shoeService *shoeServiceGrpc.Server
	ossService  ossServiceApi.OssServiceClient
	dictService *dictServiceGrpc.Server
}

func New(c *conf.Service) *Service {
	var (
		ossConn *googleGrpc.ClientConn
		err     error
	)

	service := &Service{
		shoeService: shoeServiceGrpc.New(shoeServiceService.New()),
		dictService: dictServiceGrpc.New(dictServiceService.New()),
	}

	ossConn, err = grpc.NewConnection(c.OssService)
	if err != nil {
		panic(err)
	}
	service.ossService = ossServiceApi.NewOssServiceClient(ossConn)
	//err = commonRedis.RedisClient.ScriptLoad(context.Background(), ShapeCodeListKey, updateScript)
	//if err != nil {
	//	panic(err)
	//}
	return service
}
