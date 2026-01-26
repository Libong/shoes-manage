package service

import (
	"encoding/json"
	"libong/common/context"
	commonMysql "libong/common/orm/mysql"
	commonRedis "libong/common/redis"
	"shoe-manager/app/interface/shoe/api"
	"shoe-manager/app/interface/shoe/dao"
	shoeServiceApi "shoe-manager/app/service/shoe/api"
	"shoe-manager/errors"
	ossServiceApi "shoe-manager/rpc/oss/api"
)

const updateScript = `
local key = ARGV[1]
local delta = tonumber(ARGV[2])  -- 正数表示加，负数表示减
local hash_key = 'shapeCode:stat'

local current = redis.call('HGET', hash_key, key) or 0
current = tonumber(current)
local newVal = current + delta

if newVal == 0 then
    redis.call('HDEL', hash_key, key)
    redis.call('SREM', 'shapeCodeList', key)
    return 0
else
    redis.call('HSET', hash_key, key, newVal)
    redis.call('SADD', 'shapeCodeList', key)
    return newVal
end
`

var (
	ShapeCodeListKey = "shapeCodeList"
)

func (s *Service) AddShoe(ctx context.Context, req *api.AddShoeReq) error {
	if req.ShoeSize == "" || req.Material == "" || req.ShapeCode == "" {
		return errors.ParamErrorOrEmpty
	}

	reqApi := &shoeServiceApi.AddShoeReq{
		ShapeCode: req.ShapeCode,
		Material:  req.Material,
		ShoeSize:  req.ShoeSize,
		IsHot:     req.IsHot,
		IsPresale: req.IsPresale,
	}
	if len(req.Pictures) != 0 {
		picturesBytes, err := json.Marshal(req.Pictures)
		if err != nil {
			return err
		}
		reqApi.Pictures = string(picturesBytes)
	}
	if len(req.Videos) != 0 {
		videosBytes, err := json.Marshal(req.Videos)
		if err != nil {
			return err
		}
		reqApi.Videos = string(videosBytes)
	}
	_, err := s.shoeService.AddShoe(ctx, reqApi)
	if err != nil {
		return err
	}
	//var args []interface{}
	//args = append(args, req.ShapeCode)
	//args = append(args, 1)
	//_, err = commonRedis.RedisClient.DoScript(ctx, ShapeCodeListKey, args)
	//if err != nil {
	//	return err
	//}
	return nil
}
func (s *Service) UpdateShoe(ctx context.Context, req *api.UpdateShoeReq) error {
	if req.ShoeSize == "" || req.Material == "" || req.ShapeCode == "" || req.ShoeId == "" {
		return errors.ParamErrorOrEmpty
	}

	shoeByIdResp, err := s.shoeService.ShoeById(ctx, &shoeServiceApi.ShoeByIdReq{
		Id: req.ShoeId,
	})
	if err != nil {
		return err
	}
	if shoeByIdResp.Shoe == nil {
		return errors.ShoeNotExist
	}

	reqApi := &shoeServiceApi.UpdateShoeReq{
		ShoeId:    req.ShoeId,
		ShapeCode: req.ShapeCode,
		Material:  req.Material,
		ShoeSize:  req.ShoeSize,
		IsHot:     req.IsHot,
		IsPresale: req.IsPresale,
	}
	if len(req.Pictures) != 0 {
		picturesBytes, err := json.Marshal(req.Pictures)
		if err != nil {
			return err
		}
		reqApi.Pictures = string(picturesBytes)
	}
	if len(req.Videos) != 0 {
		videosBytes, err := json.Marshal(req.Videos)
		if err != nil {
			return err
		}
		reqApi.Videos = string(videosBytes)
	}

	if req.ShapeCode != shoeByIdResp.Shoe.ShapeCode {
		var addArgs []interface{}
		addArgs = append(addArgs, req.ShapeCode)
		addArgs = append(addArgs, 1)
		_, err = commonRedis.RedisClient.DoScript(ctx, ShapeCodeListKey, addArgs)
		if err != nil {
			return err
		}
		var delArgs []interface{}
		delArgs = append(delArgs, shoeByIdResp.Shoe.ShapeCode)
		delArgs = append(delArgs, -1)
		_, err = commonRedis.RedisClient.DoScript(ctx, ShapeCodeListKey, delArgs)
		if err != nil {
			return err
		}
	}
	_, err = s.shoeService.UpdateShoe(ctx, reqApi)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteShoe(ctx context.Context, req *api.DeleteShoeReq) error {
	if req.ShoeId == "" {
		return errors.ParamErrorOrEmpty
	}
	shoeByIdResp, err := s.shoeService.ShoeById(ctx, &shoeServiceApi.ShoeByIdReq{
		Id: req.ShoeId,
	})
	if err != nil {
		return err
	}
	if shoeByIdResp.Shoe == nil {
		return errors.ShoeNotExist
	}
	_, err = s.shoeService.DeleteShoe(ctx, &shoeServiceApi.DeleteShoeReq{
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
	//var args []interface{}
	//args = append(args, shoeByIdResp.Shoe.ShoeSize)
	//args = append(args, -1)
	//_, err = commonRedis.RedisClient.DoScript(ctx, ShapeCodeListKey, args)
	//if err != nil {
	//	return err
	//}
	return nil
}
func (s *Service) SearchShoesPage(ctx context.Context, req *api.SearchShoesPageReq) (*api.SearchShoesPageResp, error) {
	if req.PageNum == 0 || req.PageSize == 0 {
		return nil, errors.ParamErrorOrEmpty
	}
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
		if req.HasFile {
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
		}
		resp.List = append(resp.List, respShoe)
	}
	if len(fileIds) != 0 && req.HasFile {
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
func (s *Service) SearchSelectList(ctx context.Context, req *api.SearchSelectListReq) (*api.SearchSelectListResp, error) {
	//members, err := commonRedis.RedisClient.SMembers(ctx, ShapeCodeListKey)
	//if err != nil {
	//	return nil, err
	//}
	resp, err := statSelectShapeCode()
	if err != nil {
		return nil, err
	}
	return &api.SearchSelectListResp{
		Map: resp,
	}, nil
}

var (
	shapeCodeStat = `
SELECT *
FROM (
    SELECT *,
           ROW_NUMBER() OVER (PARTITION BY shape_code ORDER BY created_at DESC) as rn
    FROM shoe
    WHERE deleted_at IS NULL
) t WHERE rn = 1;
`
)

func statSelectShapeCode() (map[string]*api.SelectList, error) {
	client := commonMysql.MysqlClient
	respMap := make(map[string]*api.SelectList)
	var respList []*shoeServiceApi.Shoe
	err := client.Raw(shapeCodeStat).Scan(&respList).Error
	if err != nil {
		return nil, err
	}
	for _, item := range respList {
		if respMap[item.Material] == nil {
			respMap[item.Material] = &api.SelectList{}
		}
		respMap[item.Material].Items = append(respMap[item.Material].Items, item.ShapeCode)
	}
	return respMap, nil
}
