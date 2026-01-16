package service

import (
	"libong/common/context"
	"libong/common/snowflake"
	commonTool "libong/common/tool"
	"shoe-manager/app/service/category/api"
	"shoe-manager/app/service/category/model"
)

func (s *Service) AddCategory(ctx context.Context, req *api.AddCategoryReq) (*api.AddCategoryResp, error) {
	addModel := &model.Category{}
	err := commonTool.CopyByJSON(req, addModel)
	if err != nil {
		return nil, err
	}
	addModel.AppId = ctx.AppID()
	addModel.CategoryId = snowflake.SnowflakeWorker.NextID().String()
	err = s.dao.AddCategory(ctx, addModel)
	if err != nil {
		return nil, err
	}
	return &api.AddCategoryResp{
		Id: addModel.CategoryId,
	}, nil
}

func (s *Service) UpdateCategory(ctx context.Context, req *api.UpdateCategoryReq) error {
	//先更新非零值 只有map才能更新零值
	updateModel := &model.UpdateCategoryReq{}
	err := commonTool.CopyByJSON(req, updateModel)
	if err != nil {
		return err
	}
	err = s.dao.UpdateCategory(ctx, updateModel, req.CategoryId)
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
		err = s.dao.UpdateCategory(ctx, updateMap, req.CategoryId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) DeleteCategory(ctx context.Context, req *api.DeleteCategoryReq) error {
	return s.dao.DeleteCategory(ctx, req.Id)
}

func (s *Service) CategoryByID(ctx context.Context, req *api.CategoryByIDReq) (*api.CategoryByIDResp, error) {
	m, err := s.dao.CategoryByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	respApi := &api.CategoryByIDResp{Category: &api.Category{}}
	if m == nil {
		return respApi, nil
	}
	err = commonTool.CopyByJSON(m, respApi.Category)
	if err != nil {
		return nil, err
	}
	return respApi, nil
}

func (s *Service) SearchCategoriesPage(ctx context.Context, req *api.SearchCategoriesPageReq) (*api.SearchCategoriesPageResp, error) {
	models, err := s.dao.SearchCategoriesPage(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &api.SearchCategoriesPageResp{}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respModel := &api.Category{}
		err = commonTool.CopyByJSON(m, respModel)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, respModel)
	}
	return resp, nil
}

func (s *Service) CountCategory(ctx context.Context, req *api.CountCategoryReq) (*api.CountCategoryResp, error) {
	count, err := s.dao.CountCategory(ctx, req.Condition)
	if err != nil {
		return nil, err
	}
	return &api.CountCategoryResp{
		Total: count,
	}, nil
}
func (s *Service) CategoriesByIds(ctx context.Context, req *api.CategoriesByIdsReq) (*api.CategoriesByIdsResp, error) {
	models, err := s.dao.CategoriesByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	resp := &api.CategoriesByIdsResp{
		Map: make(map[string]*api.Category),
	}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respModel := &api.Category{}
		err = commonTool.CopyByJSON(m, respModel)
		if err != nil {
			return nil, err
		}
		resp.Map[m.CategoryId] = respModel
	}
	return resp, nil
}
