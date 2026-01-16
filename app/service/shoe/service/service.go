package service

import (
	"shoe-manager/app/service/shoe/dao"
)

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	return &Service{
		dao: dao.New(),
	}
}
