package service

import (
	"libong/common/context"
	commonTool "libong/common/tool"
	"shoe-manager/app/interface/dict/api"
	dictServiceApi "shoe-manager/app/service/dict/api"
	"shoe-manager/errors"
)

func (s *Service) AddDict(ctx context.Context, req *api.AddDictReq) error {
	if req.Identifier == "" || req.Name == "" || req.DictType == "" {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.dictService.AddDict(ctx, &dictServiceApi.AddDictReq{
		Name:       req.Name,
		Identifier: req.Identifier,
		DictType:   req.DictType,
		DictValue:  req.DictValue,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UpdateDict(ctx context.Context, req *api.UpdateDictReq) error {
	if req.Identifier == "" || req.Name == "" || req.DictId == "" {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.dictService.UpdateDict(ctx, &dictServiceApi.UpdateDictReq{
		DictId:     req.DictId,
		Name:       req.Name,
		Identifier: req.Identifier,
		DictType:   req.DictType,
		DictValue:  req.DictValue,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteDict(ctx context.Context, req *api.DeleteDictReq) error {
	if req.DictId == "" {
		return errors.ParamErrorOrEmpty
	}
	_, err := s.dictService.DeleteDict(ctx, &dictServiceApi.DeleteDictReq{
		Id: req.DictId,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) SearchDictionariesPage(ctx context.Context, req *api.SearchDictionariesPageReq) (*api.
	SearchDictionariesPageResp, error) {
	resp := &api.SearchDictionariesPageResp{}
	condition := &dictServiceApi.DictCondition{
		Name:       req.Name,
		Identifier: req.Identifier,
		DictType:   req.DictType,
	}
	countDictResp, err := s.dictService.CountDict(ctx, &dictServiceApi.CountDictReq{
		Condition: condition,
	})
	if err != nil {
		return nil, err
	}
	if countDictResp.Total == 0 {
		return resp, nil
	}
	resp.Total = countDictResp.Total
	searchDictionariesPageResp, err := s.dictService.SearchDictionariesPage(ctx,
		&dictServiceApi.SearchDictionariesPageReq{
			Condition: condition,
			PageSize:  req.PageSize,
			PageNum:   req.PageNum,
		})
	if err != nil {
		return nil, err
	}

	for _, dict := range searchDictionariesPageResp.List {
		respDict := &api.Dict{
			DictId:     dict.DictId,
			Name:       dict.Name,
			Identifier: dict.Identifier,
			DictType:   dict.DictType,
			DictValue:  dict.DictValue,
		}
		resp.List = append(resp.List, respDict)
	}
	return resp, nil
}
func (s *Service) DictById(ctx context.Context, req *api.DictByIdReq) (*api.DictByIdResp, error) {
	if req.DictId == "" {
		return nil, errors.ParamErrorOrEmpty
	}
	resp := &api.DictByIdResp{}
	dictByIDResp, err := s.dictService.DictByID(ctx, &dictServiceApi.DictByIDReq{
		Id: req.DictId,
	})
	if err != nil {
		return nil, err
	}
	if dictByIDResp.Dict == nil {
		return resp, nil
	}
	err = commonTool.CopyByJSON(dictByIDResp.Dict, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
