package dao

import (
	"gorm.io/gorm"
	"libong/common/context"
	"libong/common/log"
	commonMysql "libong/common/orm/mysql"
	"shoe-manager/app/service/dict/api"
	"shoe-manager/app/service/dict/model"
)

const (
	_dictTable = "dict"
)

func (d *Dao) AddDict(ctx context.Context, dict *model.Dict) error {
	db := d.Context(ctx).Table(_dictTable).Create(&dict)
	if err := db.Error; err != nil {
		log.Error(ctx, "AddDict(%+v) error(%+v)", dict, err)
		return err
	}
	return nil
}
func (d *Dao) UpdateDict(ctx context.Context, field interface{}, dictId string) error {
	db := d.Context(ctx).Table(_dictTable).Where("dict_id = ?", dictId).Updates(field)
	if err := db.Error; err != nil {
		log.Error(ctx, "UpdateDict dictId(%v) error(%+v)", dictId, err)
		return err
	}
	return nil
}
func (d *Dao) DeleteDict(ctx context.Context, dictId string) error {
	db := d.Context(ctx).Table(_dictTable).Where("dict_id = ?", dictId).Delete(&model.Dict{})
	if err := db.Error; err != nil {
		log.Error(ctx, "DeleteDict(%+v) error(%+v)", dictId, err)
		return err
	}
	return nil
}
func (d *Dao) DictByID(ctx context.Context, dictId string) (*model.Dict, error) {
	var respModel *model.Dict
	db := d.Context(ctx).Table(_dictTable).Where("dict_id = ?", dictId)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "dictById(%+v) error(%+v)", dictId, err)
		return nil, err
	}
	return respModel, nil
}
func (d *Dao) SearchCategoriesPage(ctx context.Context, req *api.SearchDictionariesPageReq) ([]*model.Dict, error) {
	var respModel []*model.Dict
	db := d.Context(ctx, commonMysql.AppIDCondition).Table(_dictTable).Order("id desc")
	if req.PageNum != 0 && req.PageSize != 0 {
		db.Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize))
	}
	buildDictConditionDB(db, req.Condition)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchCategoriesPage(%+v) error(%+v)", req, err)
		return nil, err
	}
	return respModel, nil
}

func buildDictConditionDB(db *gorm.DB, req *api.DictCondition) {
	if len(req.DictIds) != 0 {
		db.Where("dict_ids in ?", req.DictIds)
	}

}

func (d *Dao) CountDict(ctx context.Context, req *api.DictCondition) (int64, error) {
	var count int64
	db := d.Context(ctx).Table(_dictTable).Model(&model.Dict{})
	buildDictConditionDB(db, req)
	db.Count(&count)
	if err := db.Error; err != nil {
		log.Error(ctx, "CountPage(%+v) error(%+v)", req, err)
		return 0, err
	}
	return count, nil
}
func (d *Dao) DictionariesByIds(ctx context.Context, dictIds []string) ([]*model.Dict, error) {
	var respModel []*model.Dict
	db := d.Context(ctx).Table(_dictTable).Where("dict_id in ?", dictIds)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "DictionariesByIds  error.  dictIds:(%+v);error:(%+v)", dictIds, err)
		return nil, err
	}
	return respModel, nil
}
