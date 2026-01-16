package service

import (
	"libong/common/context"
	commonTool "libong/common/tool"
	"shoe-manager/app/service/shoe/api"
	"shoe-manager/app/service/shoe/model"
)

func (s *Service) BatchAddAccountShoe(ctx context.Context, req *api.BatchAddAccountShoeReq) error {
	var models []*model.AccountShoe
	for _, item := range req.List {
		modelAccountRole := &model.AccountShoe{}
		err := commonTool.CopyByJSON(item, modelAccountRole)
		if err != nil {
			return err
		}
		models = append(models, modelAccountRole)
	}
	err := s.dao.BatchAddAccountShoe(ctx, models)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteAccountShoe(ctx context.Context, req *api.DeleteAccountShoeReq) error {
	return s.dao.DeleteAccountShoe(ctx, req.AccountIds, req.ShoeIds)
}

func (s *Service) SearchAccountShoes(ctx context.Context, req *api.SearchAccountShoesReq) (*api.SearchAccountShoesResp, error) {
	resp := &api.SearchAccountShoesResp{}
	models, err := s.dao.SearchAccountShoes(ctx, req)
	if err != nil {
		return nil, err
	}
	if models == nil {
		return resp, nil
	}
	for _, item := range models {
		respAccountRole := &api.AccountShoe{}
		err = commonTool.CopyByJSON(item, respAccountRole)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, respAccountRole)
	}
	countAccountRole, err := s.dao.CountAccountShoe(ctx, req)
	if err != nil {
		return nil, err
	}
	resp.Total = countAccountRole
	return resp, nil
}
