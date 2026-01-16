package service

import (
	"shoe-manager/app/service/category/dao"
)

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	return &Service{
		dao: dao.New(),
	}
}
