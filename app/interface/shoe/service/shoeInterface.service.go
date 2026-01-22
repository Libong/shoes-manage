package service

import (
	"encoding/json"
	"libong/common/context"
	"shoe-manager/app/interface/shoe/api"
	"shoe-manager/app/interface/shoe/dao"
	shoeServiceApi "shoe-manager/app/service/shoe/api"
	"shoe-manager/errors"
	ossServiceApi "shoe-manager/rpc/oss/api"
)

func (s *Service) AddShoe(ctx context.Context, req *api.AddShoeReq) error {
	if req.ShoeSize == "" || req.Material == "" || req.ShapeCode == "" {
		return errors.ParamErrorOrEmpty
	}
	picturesBytes, err := json.Marshal(req.Pictures)
	if err != nil {
		return err
	}
	videosBytes, err := json.Marshal(req.Videos)
	if err != nil {
		return err
	}
	_, err = s.shoeService.AddShoe(ctx, &shoeServiceApi.AddShoeReq{
		ShapeCode: req.ShapeCode,
		Material:  req.Material,
		ShoeSize:  req.ShoeSize,
		Pictures:  string(picturesBytes),
		Videos:    string(videosBytes),
		IsHot:     req.IsHot,
		IsPresale: req.IsPresale,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UpdateShoe(ctx context.Context, req *api.UpdateShoeReq) error {
	if req.ShoeSize == "" || req.Material == "" || req.ShapeCode == "" || req.ShoeId == "" {
		return errors.ParamErrorOrEmpty
	}
	picturesBytes, err := json.Marshal(req.Pictures)
	if err != nil {
		return err
	}
	videosBytes, err := json.Marshal(req.Videos)
	if err != nil {
		return err
	}
	_, err = s.shoeService.UpdateShoe(ctx, &shoeServiceApi.UpdateShoeReq{
		ShoeId:    req.ShoeId,
		ShapeCode: req.ShapeCode,
		Material:  req.Material,
		ShoeSize:  req.ShoeSize,
		Pictures:  string(picturesBytes),
		Videos:    string(videosBytes),
		IsHot:     req.IsHot,
		IsPresale: req.IsPresale,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteShoe(ctx context.Context, req *api.DeleteShoeReq) error {
	if req.ShoeId == "" {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.shoeService.DeleteShoe(ctx, &shoeServiceApi.DeleteShoeReq{
		Ids: []string{req.ShoeId},
	})
	if err != nil {
		return err
	}
	_, err = s.shoeService.DeleteAccountShoe(ctx, &shoeServiceApi.DeleteAccountShoeReq{
		ShoeIds: []string{req.ShoeId},
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) SearchShoesPage(ctx context.Context, req *api.SearchShoesPageReq) (*api.SearchShoesPageResp, error) {
	resp := &api.SearchShoesPageResp{}
	var searchShoesPageResp *shoeServiceApi.SearchShoesPageResp
	var err error
	favourShoeMap := make(map[string]int)
	if !req.ByFavour {
		condition := &shoeServiceApi.SearchShoeCondition{
			ShapeCode: req.ShapeCode,
			Material:  req.Material,
			ShoeSize:  req.ShoeSize,
			IsPresale: req.IsPresale,
			IsHot:     req.IsHot,
		}
		countShoesResp, err := s.shoeService.CountShoe(ctx, &shoeServiceApi.CountShoesReq{
			Condition: condition,
		})
		if err != nil {
			return nil, err
		}
		if countShoesResp.Total == 0 {
			return resp, nil
		}
		resp.Total = countShoesResp.Total
		searchShoesPageResp, err = s.shoeService.SearchShoesPage(ctx, &shoeServiceApi.SearchShoesPageReq{
			Condition: condition,
			PageSize:  req.PageSize,
			PageNum:   req.PageNum,
		})
		if err != nil {
			return nil, err
		}
		var shoeIds []string
		for _, shoe := range searchShoesPageResp.List {
			shoeIds = append(shoeIds, shoe.ShoeId)
		}
		searchAccountShoesResp, err := s.shoeService.SearchAccountShoes(ctx, &shoeServiceApi.SearchAccountShoesReq{
			AccountIds: []string{ctx.User().UID},
			ShoeIds:    shoeIds,
		})
		if err != nil {
			return nil, err
		}
		for _, accountShoe := range searchAccountShoesResp.List {
			favourShoeMap[accountShoe.ShoeId]++
		}
	} else {
		searchShoesPageResp, resp.Total, err = dao.SearchAccountShoeJoin(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	var fileIds []string
	for _, shoe := range searchShoesPageResp.List {
		respShoe := &api.Shoe{
			ShoeId:    shoe.ShoeId,
			ShapeCode: shoe.ShapeCode,
			Material:  shoe.Material,
			ShoeSize:  shoe.ShoeSize,
			IsHot:     shoe.IsHot,
			IsPresale: shoe.IsPresale,
			IsFavour:  true,
		}
		if !req.ByFavour && favourShoeMap[shoe.ShoeId] == 0 {
			respShoe.IsFavour = false
		}
		if shoe.Pictures != "" {
			err = json.Unmarshal([]byte(shoe.Pictures), &respShoe.Pictures)
			if err != nil {
				return nil, err
			}
		}
		if shoe.Videos != "" {
			err = json.Unmarshal([]byte(shoe.Videos), &respShoe.Videos)
			if err != nil {
				return nil, err
			}
		}
		for _, picture := range respShoe.Pictures {
			fileIds = append(fileIds, picture.Id)
		}
		for _, video := range respShoe.Videos {
			fileIds = append(fileIds, video.Id)
		}
		resp.List = append(resp.List, respShoe)
	}
	if len(fileIds) != 0 {
		makeFileUrlResp, err := s.ossService.MakeFileUrl(ctx, &ossServiceApi.MakeFileUrlReq{
			Keys: fileIds,
		})
		if err != nil {
			return nil, err
		}
		fileUrlMap := makeFileUrlResp.Map
		for _, shoe := range resp.List {
			for _, picture := range shoe.Pictures {
				picture.Url = fileUrlMap[picture.Id]
			}
			for _, video := range shoe.Videos {
				video.Url = fileUrlMap[video.Id]
			}
		}
	}
	return resp, nil
}
func (s *Service) ShoeById(ctx context.Context, req *api.ShoeByIdReq) (*api.ShoeByIdResp, error) {
	if req.ShoeId == "" {
		return nil, errors.ParamErrorOrEmpty
	}
	if ctx.User() == nil {
		return nil, errors.AccountNotExist
	}
	shoeByIdResp, err := s.shoeService.ShoeById(ctx, &shoeServiceApi.ShoeByIdReq{
		Id: req.ShoeId,
	})
	if err != nil {
		return nil, err
	}
	resp := &api.ShoeByIdResp{}
	shoe := shoeByIdResp.Shoe
	if shoe == nil {
		return resp, nil
	}
	resp = &api.ShoeByIdResp{
		ShoeId:    shoe.ShoeId,
		ShapeCode: shoe.ShapeCode,
		Material:  shoe.Material,
		ShoeSize:  shoe.ShoeSize,
		Pictures:  nil,
		Videos:    nil,
		IsHot:     shoe.IsHot,
		IsPresale: shoe.IsPresale,
	}

	searchAccountShoesResp, err := s.shoeService.SearchAccountShoes(ctx, &shoeServiceApi.SearchAccountShoesReq{
		AccountIds: []string{ctx.User().UID},
		ShoeIds:    []string{req.ShoeId},
	})
	if err != nil {
		return nil, err
	}
	if len(searchAccountShoesResp.List) != 0 {
		resp.IsFavour = true
	}
	if shoe.Pictures != "" {
		err = json.Unmarshal([]byte(shoe.Pictures), &resp.Pictures)
		if err != nil {
			return nil, err
		}
	}
	if shoe.Videos != "" {
		err = json.Unmarshal([]byte(shoe.Videos), &resp.Videos)
		if err != nil {
			return nil, err
		}
	}

	var fileIds []string
	for _, picture := range resp.Pictures {
		fileIds = append(fileIds, picture.Id)
	}
	for _, video := range resp.Videos {
		fileIds = append(fileIds, video.Id)
	}

	if len(fileIds) != 0 {
		makeFileUrlResp, err := s.ossService.MakeFileUrl(ctx, &ossServiceApi.MakeFileUrlReq{
			Keys: fileIds,
		})
		if err != nil {
			return nil, err
		}
		fileUrlMap := makeFileUrlResp.Map
		for _, picture := range resp.Pictures {
			picture.Url = fileUrlMap[picture.Id]
		}
		for _, video := range resp.Videos {
			video.Url = fileUrlMap[video.Id]
		}
	}

	return resp, nil
}
func (s *Service) UpdateShoeHot(ctx context.Context, req *api.UpdateShoeHotReq) error {
	if len(req.ShoeIds) == 0 {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.shoeService.BatchUpdateShoe(ctx, &shoeServiceApi.BatchUpdateShoeReq{
		Ids:   req.ShoeIds,
		IsHot: req.IsHot,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) ChangeShoeFavour(ctx context.Context, req *api.ChangeShoeFavourReq) error {
	if len(req.ShoeIds) == 0 {
		return errors.ParamErrorOrEmpty
	}
	if ctx.User() == nil {
		return errors.AccountNotExist
	}
	uid := ctx.User().UID
	if !req.IsFavour {
		_, err := s.shoeService.DeleteAccountShoe(ctx, &shoeServiceApi.DeleteAccountShoeReq{
			AccountIds: []string{uid},
			ShoeIds:    req.ShoeIds,
		})
		if err != nil {
			return err
		}
		return nil
	}
	existShoeFavourMap := make(map[string]uint32)
	searchAccountShoesResp, err := s.shoeService.SearchAccountShoes(ctx, &shoeServiceApi.SearchAccountShoesReq{
		AccountIds: []string{uid},
		ShoeIds:    req.ShoeIds,
	})
	if err != nil {
		return err
	}
	for _, accountShoe := range searchAccountShoesResp.List {
		existShoeFavourMap[accountShoe.ShoeId]++
	}
	var addAccountShoes []*shoeServiceApi.AccountShoe
	for _, changeShoeId := range req.ShoeIds {
		if existShoeFavourMap[changeShoeId] == 0 {
			addAccountShoes = append(addAccountShoes, &shoeServiceApi.AccountShoe{
				AccountId: uid,
				ShoeId:    changeShoeId,
			})
		}
	}
	if len(addAccountShoes) != 0 {
		_, err = s.shoeService.BatchAddAccountShoe(ctx, &shoeServiceApi.BatchAddAccountShoeReq{
			List: addAccountShoes,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
