package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	qContext "libong/common/context"
	"shoe-manager/app/service/shoe/api"
)

func (s *Server) AddShoe(ctx context.Context, req *api.AddShoeReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.AddShoe(ctx.(qContext.Context), req)
}

func (s *Server) DeleteShoe(ctx context.Context, req *api.DeleteShoeReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.DeleteShoe(ctx.(qContext.Context), req)
}

func (s *Server) UpdateShoe(ctx context.Context, req *api.UpdateShoeReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.UpdateShoe(ctx.(qContext.Context), req)
}

func (s *Server) SearchShoesPage(ctx context.Context, req *api.SearchShoesPageReq) (*api.SearchShoesPageResp, error) {
	return s.service.SearchShoesPage(ctx.(qContext.Context), req)
}

func (s *Server) CountShoe(ctx context.Context, req *api.CountShoesReq) (*api.CountShoesResp, error) {
	return s.service.CountShoe(ctx.(qContext.Context), req)
}

func (s *Server) ShoeById(ctx context.Context, req *api.ShoeByIdReq) (*api.ShoeByIdResp, error) {
	return s.service.ShoeById(ctx.(qContext.Context), req)
}

func (s *Server) ShoesByIds(ctx context.Context, req *api.ShoesByIdsReq) (*api.ShoesByIdsResp, error) {
	return s.service.ShoesByIds(ctx.(qContext.Context), req)
}
func (s *Server) BatchUpdateShoe(ctx context.Context, req *api.BatchUpdateShoeReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.BatchUpdateShoe(ctx.(qContext.Context), req)
}
func (s *Server) BatchAddAccountShoe(ctx context.Context, req *api.BatchAddAccountShoeReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.BatchAddAccountShoe(ctx.(qContext.Context), req)
}

func (s *Server) DeleteAccountShoe(ctx context.Context, req *api.DeleteAccountShoeReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.DeleteAccountShoe(ctx.(qContext.Context), req)
}

func (s *Server) SearchAccountShoes(ctx context.Context, req *api.SearchAccountShoesReq) (*api.SearchAccountShoesResp, error) {
	return s.service.SearchAccountShoes(ctx.(qContext.Context), req)
}
