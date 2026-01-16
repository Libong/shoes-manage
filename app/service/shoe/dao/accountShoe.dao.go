package dao

import (
	"libong/common/context"
	"libong/common/log"
	"shoe-manager/app/service/shoe/api"
	"shoe-manager/app/service/shoe/model"
)

const (
	_AccountShoeTable = "account_shoe"
)

func (d *Dao) BatchAddAccountShoe(ctx context.Context, accountShoes []*model.AccountShoe) error {
	db := d.Context(ctx).Table(_AccountShoeTable).CreateInBatches(accountShoes, 10)
	if err := db.Error; err != nil {
		log.Error(ctx, "BatchAddAccountShoe  error.  accountShoes:(%+v);error:(%+v)", accountShoes, err)
		return err
	}
	return nil
}

func (d *Dao) DeleteAccountShoe(ctx context.Context, accountIds []string, shoeIds []string) error {
	db := d.Context(ctx).Table(_AccountShoeTable)
	if len(accountIds) != 0 {
		db.Where("account_id in ?", accountIds)
	}
	if len(shoeIds) != 0 {
		db.Where("shoe_id in ?", shoeIds)
	}
	db.Delete(&model.AccountShoe{})
	if err := db.Error; err != nil {
		log.Error(ctx, "DeleteAccountShoe  error.  accountIds:(%+v);shoeIds:(%+v);error:(%+v)", accountIds, shoeIds, err)
		return err
	}
	return nil
}
func (d *Dao) SearchAccountShoes(ctx context.Context, req *api.SearchAccountShoesReq) ([]*model.AccountShoe,
	error) {
	var respModels []*model.AccountShoe
	db := d.Context(ctx).Table(_AccountShoeTable)
	if len(req.AccountIds) != 0 {
		db.Where("account_id in ?", req.AccountIds)
	}
	if len(req.ShoeIds) != 0 {
		db.Where("shoe_id in ?", req.ShoeIds)
	}
	db.Find(&respModels)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchAccountShoes  error.  req:(%+v);error:(%+v)", req, err)
		return nil, err
	}
	return respModels, nil
}
func (d *Dao) CountAccountShoe(ctx context.Context, req *api.SearchAccountShoesReq) (int64, error) {
	var total int64
	db := d.Context(ctx).Table(_AccountShoeTable)
	if len(req.AccountIds) != 0 {
		db.Where("account_id in ?", req.AccountIds)
	}
	if len(req.ShoeIds) != 0 {
		db.Where("shoe_id in ?", req.ShoeIds)
	}
	db.Count(&total)
	if err := db.Error; err != nil {
		log.Error(ctx, "CountAccountShoe  error.  req:(%+v);error:(%+v)", req, err)
		return 0, err
	}
	return total, nil
}
