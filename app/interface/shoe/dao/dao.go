package dao

import (
	"libong/common/context"
	commonMysql "libong/common/orm/mysql"
	"shoe-manager/app/interface/shoe/api"
	serviceApi "shoe-manager/app/service/shoe/api"
	"shoe-manager/errors"
)

var (
	sql = `
SELECT s.*
FROM shoe s
INNER JOIN (
    -- 先分页获取用户收藏的物品ID
    SELECT shoe_id
    FROM account_shoe
    WHERE account_id = ?  -- 当前用户ID
    LIMIT ?, ?  -- 分页参数
) AS rel ON s.shoe_id = rel.shoe_id
WHERE 1=1 
`
	countSql = `
SELECT count(*)
FROM shoe s
INNER JOIN (
    -- 先分页获取用户收藏的物品ID
    SELECT shoe_id
    FROM account_shoe
    WHERE account_id = ?  -- 当前用户ID
) AS rel ON s.shoe_id = rel.shoe_id
WHERE 1=1 
`
	materialCondition  = ` and s.material like ?`
	shoeSizeCondition  = ` and s.shoe_size like ?`
	shapeCodeCondition = ` and s.shape_code like ?`
	isHotCondition     = ` and s.is_hot = ?`
	isPresaleCondition = ` and s.is_presale = ?`
	finalSql           = ` ORDER BY s.created_at DESC;`
	finalTag           = `;`
)

func SearchAccountShoeJoin(ctx context.Context, req *api.SearchShoesPageReq) (*serviceApi.SearchShoesPageResp,
	int64, error) {
	if ctx.User() == nil {
		return nil, 0, errors.AccountNotExist
	}
	uid := ctx.User().UID
	client := commonMysql.MysqlClient
	originSql := sql
	originCountSql := countSql
	var args []interface{}
	var countArgs []interface{}
	args = append(args, uid)
	countArgs = append(countArgs, uid)
	offset := (req.PageNum - 1) * req.PageSize
	args = append(args, offset)
	args = append(args, offset+req.PageSize)
	if req.Material != "" {
		originSql += materialCondition
		originCountSql += materialCondition
		args = append(args, "%"+req.Material+"%")
		countArgs = append(countArgs, "%"+req.Material+"%")
	}
	if req.ShoeSize != "" {
		originSql += shoeSizeCondition
		originCountSql += shoeSizeCondition
		args = append(args, "%"+req.ShoeSize+"%")
		countArgs = append(countArgs, "%"+req.ShoeSize+"%")
	}
	if req.ShapeCode != "" {
		originSql += shapeCodeCondition
		originCountSql += shapeCodeCondition
		args = append(args, "%"+req.ShapeCode+"%")
		countArgs = append(countArgs, "%"+req.ShapeCode+"%")
	}
	if req.IsHot != 0 {
		originSql += isHotCondition
		originCountSql += isHotCondition
		args = append(args, req.IsHot)
		countArgs = append(countArgs, req.IsHot)
	}
	if req.IsPresale != 0 {
		originSql += isPresaleCondition
		originCountSql += isPresaleCondition
		args = append(args, req.IsPresale)
		countArgs = append(countArgs, req.IsPresale)
	}

	originSql += finalSql
	originCountSql += finalTag
	resp := &serviceApi.SearchShoesPageResp{}
	var total int64
	row := client.Context(ctx).Raw(originCountSql, countArgs...).Row()
	err := row.Scan(&total)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return resp, total, nil
	}
	err = client.Context(ctx).Raw(originSql, args...).Scan(&resp.List).Error
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
