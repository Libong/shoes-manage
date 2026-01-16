package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	qContext "libong/common/context"
	"shoe-manager/app/service/category/api"
)

func (s *Server) AddCategory(ctx context.Context, req *api.AddCategoryReq) (*api.AddCategoryResp, error) {
	return s.service.AddCategory(ctx.(qContext.Context), req)
}

func (s *Server) UpdateCategory(ctx context.Context, req *api.UpdateCategoryReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.UpdateCategory(ctx.(qContext.Context), req)
}

func (s *Server) DeleteCategory(ctx context.Context, req *api.DeleteCategoryReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.DeleteCategory(ctx.(qContext.Context), req)
}

func (s *Server) CategoryByID(ctx context.Context, req *api.CategoryByIDReq) (*api.CategoryByIDResp, error) {
	return s.service.CategoryByID(ctx.(qContext.Context), req)
}

func (s *Server) SearchCategoriesPage(ctx context.Context, req *api.SearchCategoriesPageReq) (*api.SearchCategoriesPageResp, error) {
	return s.service.SearchCategoriesPage(ctx.(qContext.Context), req)
}

func (s *Server) CountCategory(ctx context.Context, req *api.CountCategoryReq) (*api.CountCategoryResp, error) {
	return s.service.CountCategory(ctx.(qContext.Context), req)
}

func (s *Server) CategoriesByIds(ctx context.Context, req *api.CategoriesByIdsReq) (*api.CategoriesByIdsResp, error) {
	return s.service.CategoriesByIds(ctx.(qContext.Context), req)
}
