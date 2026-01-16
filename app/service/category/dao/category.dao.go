package dao

import (
	"gorm.io/gorm"
	"libong/common/context"
	"libong/common/log"
	commonMysql "libong/common/orm/mysql"
	"shoe-manager/app/service/category/api"
	"shoe-manager/app/service/category/model"
)

const (
	_categoryTable = "category"
)

func (d *Dao) AddCategory(ctx context.Context, category *model.Category) error {
	db := d.Context(ctx).Table(_categoryTable).Create(&category)
	if err := db.Error; err != nil {
		log.Error(ctx, "AddCategory(%+v) error(%+v)", category, err)
		return err
	}
	return nil
}
func (d *Dao) UpdateCategory(ctx context.Context, field interface{}, categoryId string) error {
	db := d.Context(ctx).Table(_categoryTable).Where("category_id = ?", categoryId).Updates(field)
	if err := db.Error; err != nil {
		log.Error(ctx, "UpdateCategory categoryId(%v) error(%+v)", categoryId, err)
		return err
	}
	return nil
}
func (d *Dao) DeleteCategory(ctx context.Context, categoryId string) error {
	db := d.Context(ctx).Table(_categoryTable).Where("category_id = ?", categoryId).Delete(&model.Category{})
	if err := db.Error; err != nil {
		log.Error(ctx, "DeleteCategory(%+v) error(%+v)", categoryId, err)
		return err
	}
	return nil
}
func (d *Dao) CategoryByID(ctx context.Context, categoryId string) (*model.Category, error) {
	var respModel *model.Category
	db := d.Context(ctx).Table(_categoryTable).Where("category_id = ?", categoryId)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "CategoryById(%+v) error(%+v)", categoryId, err)
		return nil, err
	}
	return respModel, nil
}
func (d *Dao) SearchCategoriesPage(ctx context.Context, req *api.SearchCategoriesPageReq) ([]*model.Category, error) {
	var respModel []*model.Category
	db := d.Context(ctx, commonMysql.AppIDCondition).Table(_categoryTable).Order("id desc")
	if req.PageNum != 0 && req.PageSize != 0 {
		db.Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize))
	}
	buildCategoryConditionDB(db, req.Condition)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchCategoriesPage(%+v) error(%+v)", req, err)
		return nil, err
	}
	return respModel, nil
}

func buildCategoryConditionDB(db *gorm.DB, req *api.CategoryCondition) {
	if len(req.CategoryIds) != 0 {
		db.Where("category_ids in ?", req.CategoryIds)
	}

}

func (d *Dao) CountCategory(ctx context.Context, req *api.CategoryCondition) (int64, error) {
	var count int64
	db := d.Context(ctx).Table(_categoryTable).Model(&model.Category{})
	buildCategoryConditionDB(db, req)
	db.Count(&count)
	if err := db.Error; err != nil {
		log.Error(ctx, "CountPage(%+v) error(%+v)", req, err)
		return 0, err
	}
	return count, nil
}
func (d *Dao) CategoriesByIds(ctx context.Context, categoryIds []string) ([]*model.Category, error) {
	var respModel []*model.Category
	db := d.Context(ctx).Table(_categoryTable).Where("category_id in ?", categoryIds)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "CategoriesByIds  error.  categoryIds:(%+v);error:(%+v)", categoryIds, err)
		return nil, err
	}
	return respModel, nil
}
