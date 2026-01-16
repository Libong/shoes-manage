package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	qContext "libong/common/context"
	"shoe-manager/app/service/dict/api"
)

func (s *Server) AddDict(ctx context.Context, req *api.AddDictReq) (*api.AddDictResp, error) {
	return s.service.AddDict(ctx.(qContext.Context), req)
}

func (s *Server) UpdateDict(ctx context.Context, req *api.UpdateDictReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.UpdateDict(ctx.(qContext.Context), req)
}

func (s *Server) DeleteDict(ctx context.Context, req *api.DeleteDictReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.DeleteDict(ctx.(qContext.Context), req)
}

func (s *Server) DictByID(ctx context.Context, req *api.DictByIDReq) (*api.DictByIDResp, error) {
	return s.service.DictByID(ctx.(qContext.Context), req)
}

func (s *Server) SearchCategoriesPage(ctx context.Context, req *api.SearchDictionariesPageReq) (*api.SearchDictionariesPageResp, error) {
	return s.service.SearchCategoriesPage(ctx.(qContext.Context), req)
}

func (s *Server) CountDict(ctx context.Context, req *api.CountDictReq) (*api.CountDictResp, error) {
	return s.service.CountDict(ctx.(qContext.Context), req)
}

func (s *Server) CategoriesByIds(ctx context.Context, req *api.DictionariesByIdsReq) (*api.DictionariesByIdsResp, error) {
	return s.service.CategoriesByIds(ctx.(qContext.Context), req)
}
