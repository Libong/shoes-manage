package service

import (
	"libong/common/context"
	commonTool "libong/common/tool"
	"shoe-manager/app/interface/category/api"
	categoryServiceApi "shoe-manager/app/service/category/api"
	"shoe-manager/errors"
)

func (s *Service) AddCategory(ctx context.Context, req *api.AddCategoryReq) error {
	if req.Identifier == "" || req.Name == "" {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.categoryService.AddCategory(ctx, &categoryServiceApi.AddCategoryReq{
		Name:       req.Name,
		Identifier: req.Identifier,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UpdateCategory(ctx context.Context, req *api.UpdateCategoryReq) error {
	if req.Identifier == "" || req.Name == "" || req.CategoryId == "" {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.categoryService.UpdateCategory(ctx, &categoryServiceApi.UpdateCategoryReq{
		CategoryId: req.CategoryId,
		Name:       req.Name,
		Identifier: req.Identifier,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteCategory(ctx context.Context, req *api.DeleteCategoryReq) error {
	if req.CategoryId == "" {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.categoryService.DeleteCategory(ctx, &categoryServiceApi.DeleteCategoryReq{
		Id: req.CategoryId,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) SearchCategoriesPage(ctx context.Context, req *api.SearchCategoriesPageReq) (*api.
	SearchCategoriesPageResp, error) {
	resp := &api.SearchCategoriesPageResp{}
	condition := &categoryServiceApi.CategoryCondition{
		Name:       req.Name,
		Identifier: req.Identifier,
	}
	countCategoriesResp, err := s.categoryService.CountCategory(ctx, &categoryServiceApi.CountCategoryReq{
		Condition: condition,
	})
	if err != nil {
		return nil, err
	}
	if countCategoriesResp.Total == 0 {
		return resp, nil
	}
	resp.Total = countCategoriesResp.Total
	searchCategoriesPageResp, err := s.categoryService.SearchCategoriesPage(ctx,
		&categoryServiceApi.SearchCategoriesPageReq{
			Condition: condition,
			PageSize:  req.PageSize,
			PageNum:   req.PageNum,
		})
	if err != nil {
		return nil, err
	}

	for _, category := range searchCategoriesPageResp.List {
		respCategory := &api.Category{
			CategoryId: category.CategoryId,
			Name:       category.Name,
			Identifier: category.Identifier,
		}
		resp.List = append(resp.List, respCategory)
	}
	return resp, nil
}
func (s *Service) CategoryById(ctx context.Context, req *api.CategoryByIdReq) (*api.CategoryByIdResp, error) {
	if req.CategoryId == "" {
		return nil, errors.ParamErrorOrEmpty
	}
	resp := &api.CategoryByIdResp{}
	categoryByIDResp, err := s.categoryService.CategoryByID(ctx, &categoryServiceApi.CategoryByIDReq{
		Id: req.CategoryId,
	})
	if err != nil {
		return nil, err
	}
	if categoryByIDResp.Category == nil {
		return resp, nil
	}
	err = commonTool.CopyByJSON(categoryByIDResp.Category, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
