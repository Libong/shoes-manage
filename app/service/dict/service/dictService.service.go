package service

import (
	"libong/common/context"
	"libong/common/snowflake"
	commonTool "libong/common/tool"
	"shoe-manager/app/service/dict/api"
	"shoe-manager/app/service/dict/model"
)

func (s *Service) AddDict(ctx context.Context, req *api.AddDictReq) (*api.AddDictResp, error) {
	addModel := &model.Dict{}
	err := commonTool.CopyByJSON(req, addModel)
	if err != nil {
		return nil, err
	}
	addModel.AppId = ctx.AppID()
	addModel.DictId = snowflake.SnowflakeWorker.NextID().String()
	err = s.dao.AddDict(ctx, addModel)
	if err != nil {
		return nil, err
	}
	return &api.AddDictResp{
		Id: addModel.DictId,
	}, nil
}

func (s *Service) UpdateDict(ctx context.Context, req *api.UpdateDictReq) error {
	//先更新非零值 只有map才能更新零值
	updateModel := &model.UpdateDictReq{}
	err := commonTool.CopyByJSON(req, updateModel)
	if err != nil {
		return err
	}
	err = s.dao.UpdateDict(ctx, updateModel, req.DictId)
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
		err = s.dao.UpdateDict(ctx, updateMap, req.DictId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) DeleteDict(ctx context.Context, req *api.DeleteDictReq) error {
	return s.dao.DeleteDict(ctx, req.Id)
}

func (s *Service) DictByID(ctx context.Context, req *api.DictByIDReq) (*api.DictByIDResp, error) {
	m, err := s.dao.DictByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	respApi := &api.DictByIDResp{Dict: &api.Dict{}}
	if m == nil {
		return respApi, nil
	}
	err = commonTool.CopyByJSON(m, respApi.Dict)
	if err != nil {
		return nil, err
	}
	return respApi, nil
}

func (s *Service) SearchCategoriesPage(ctx context.Context, req *api.SearchDictionariesPageReq) (*api.SearchDictionariesPageResp, error) {
	models, err := s.dao.SearchCategoriesPage(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &api.SearchDictionariesPageResp{}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respModel := &api.Dict{}
		err = commonTool.CopyByJSON(m, respModel)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, respModel)
	}
	return resp, nil
}

func (s *Service) CountDict(ctx context.Context, req *api.CountDictReq) (*api.CountDictResp, error) {
	count, err := s.dao.CountDict(ctx, req.Condition)
	if err != nil {
		return nil, err
	}
	return &api.CountDictResp{
		Total: count,
	}, nil
}
func (s *Service) CategoriesByIds(ctx context.Context, req *api.DictionariesByIdsReq) (*api.DictionariesByIdsResp, error) {
	models, err := s.dao.DictionariesByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	resp := &api.DictionariesByIdsResp{
		Map: make(map[string]*api.Dict),
	}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respModel := &api.Dict{}
		err = commonTool.CopyByJSON(m, respModel)
		if err != nil {
			return nil, err
		}
		resp.Map[m.DictId] = respModel
	}
	return resp, nil
}
