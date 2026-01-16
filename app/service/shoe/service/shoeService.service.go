package service

import (
	"libong/common/context"
	"libong/common/snowflake"
	commonTool "libong/common/tool"
	"shoe-manager/app/service/shoe/api"
	"shoe-manager/app/service/shoe/model"
)

func (s *Service) AddShoe(ctx context.Context, req *api.AddShoeReq) error {
	addModel := &model.Shoe{}
	err := commonTool.CopyByJSON(req, addModel)
	if err != nil {
		return err
	}
	addModel.ShoeId = snowflake.SnowflakeWorker.NextID().String()
	err = s.dao.AddShoe(ctx, addModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteShoe(ctx context.Context, req *api.DeleteShoeReq) error {
	return s.dao.DeleteShoe(ctx, req)
}

func (s *Service) UpdateShoe(ctx context.Context, req *api.UpdateShoeReq) error {
	//先更新非零值
	updateModel := &model.UpdateShoeReq{}
	err := commonTool.CopyByJSON(req, updateModel)
	if err != nil {
		return err
	}
	err = s.dao.UpdateShoe(ctx, updateModel, req.ShoeId)
	if err != nil {
		return err
	}
	//再更新零值
	if req.HandleZero {
		updateMap := make(map[string]interface{})
		if req.ZeroFields == nil || len(req.ZeroFields) == 0 {
			return nil
		}
		valueMap := commonTool.StructToFilterEmptyMap(*updateModel)
		zeroMap := commonTool.StructToDefaultValueMap(updateModel)
		for _, field := range req.ZeroFields {
			if zeroMap[field] != nil && valueMap[field] != nil {
				updateMap[field] = zeroMap[field]
			}
		}
		err = s.dao.UpdateShoe(ctx, updateMap, req.ShoeId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) SearchShoesPage(ctx context.Context, req *api.SearchShoesPageReq) (*api.SearchShoesPageResp, error) {
	models, err := s.dao.SearchShoesPage(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &api.SearchShoesPageResp{}
	if models == nil {
		return resp, nil
	}
	for _, item := range models {
		respShoe := &api.Shoe{}
		err = commonTool.CopyByJSON(item, respShoe)
		if err != nil {
			return nil, err
		}
		respShoe.EstablishAt = item.CreatedAt.Unix()
		respShoe.ModifyAt = item.UpdatedAt.Unix()
		resp.List = append(resp.List, respShoe)
	}
	return resp, nil
}

func (s *Service) CountShoe(ctx context.Context, req *api.CountShoesReq) (*api.CountShoesResp, error) {
	count, err := s.dao.CountShoe(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.CountShoesResp{
		Total: count,
	}, nil
}

func (s *Service) ShoeById(ctx context.Context, req *api.ShoeByIdReq) (*api.ShoeByIdResp, error) {
	respModel, err := s.dao.ShoeByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := &api.ShoeByIdResp{}
	if respModel.ID == 0 {
		return resp, nil
	}
	respShoe := &api.Shoe{}
	err = commonTool.CopyByJSON(respModel, respShoe)
	if err != nil {
		return nil, err
	}
	respShoe.EstablishAt = respModel.CreatedAt.Unix()
	respShoe.ModifyAt = respModel.UpdatedAt.Unix()
	resp.Shoe = respShoe
	return resp, nil
}

func (s *Service) ShoesByIds(ctx context.Context, req *api.ShoesByIdsReq) (*api.ShoesByIdsResp, error) {
	models, err := s.dao.ShoesByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	resp := &api.ShoesByIdsResp{
		Map: make(map[string]*api.Shoe),
	}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respShoe := &api.Shoe{}
		err = commonTool.CopyByJSON(m, respShoe)
		if err != nil {
			return nil, err
		}
		respShoe.EstablishAt = m.CreatedAt.Unix()
		respShoe.ModifyAt = m.UpdatedAt.Unix()
		resp.Map[m.ShoeId] = respShoe
	}
	return resp, nil
}
func (s *Service) BatchUpdateShoe(ctx context.Context, req *api.BatchUpdateShoeReq) error {
	updateFields := make(map[string]interface{})
	updateFields["is_hot"] = req.IsHot
	return s.dao.BatchUpdateShoe(ctx, updateFields, req.Ids)
}
