package dao

import (
	"gorm.io/gorm"
	"libong/common/context"
	"libong/common/log"
	"libong/common/orm/mysql"
	commonTool "libong/common/tool"
	"shoe-manager/app/service/shoe/api"
	"shoe-manager/app/service/shoe/model"
	"strconv"
)

/*
dao层的作用就是初始化数据库对象 orm.NewClient
还有操作数据库
*/

type Dao struct {
	*commonMysql.Client
}

func New() *Dao {
	if commonMysql.MysqlClient == nil {
		panic("MysqlClient is nil")
	}
	return &Dao{
		Client: commonMysql.MysqlClient,
	}
}

const (
	_ShoeTable = "shoe"
)

func (d *Dao) AddShoe(ctx context.Context, shoe *model.Shoe) error {
	db := d.Context(ctx).Table(_ShoeTable).Create(&shoe)
	if err := db.Error; err != nil {
		log.Error(ctx, "AddShoe  error.  shoe:(%+v);error:(%+v)", shoe, err)
		return err
	}
	return nil
}

func (d *Dao) UpdateShoe(ctx context.Context, updateFields interface{}, shoeId string) error {
	db := d.Context(ctx).Table(_ShoeTable).Where("shoe_id = ?", shoeId).Updates(updateFields)
	if err := db.Error; err != nil {
		log.Error(ctx, "UpdateShoe  error.  updateFields:(%+v);error:(%+v)", updateFields, err)
		return err
	}
	return nil
}

func (d *Dao) BatchUpdateShoe(ctx context.Context, updateFields map[string]interface{}, shoeIds []string) error {
	db := d.Context(ctx).Table(_ShoeTable).Where("shoe_id in ?", shoeIds)
	for k, v := range updateFields {
		db.Update(k, v)
	}
	if err := db.Error; err != nil {
		log.Error(ctx, "BatchUpdateShoe  error.  updateFields:(%+v);error:(%+v)", updateFields, err)
		return err
	}
	return nil
}

func (d *Dao) ShoeByID(ctx context.Context, shoeId string) (*model.Shoe, error) {
	var respModel *model.Shoe
	db := d.Context(ctx).Table(_ShoeTable).Where("shoe_id = ?", shoeId)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "ShoeByID  error.  shoeId:(%+v);error:(%+v)", shoeId, err)
		return nil, err
	}
	return respModel, nil
}
func (d *Dao) ShoeByName(ctx context.Context, name string) (*model.Shoe, error) {
	var respModel *model.Shoe
	db := d.Context(ctx).Table(_ShoeTable).Where("name = ?", name)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "ShoeByName  error.  name:(%+v);error:(%+v)", name, err)
		return nil, err
	}
	return respModel, nil
}
func (d *Dao) ShoesByIds(ctx context.Context, shoeIds []string) ([]*model.Shoe, error) {
	var respModel []*model.Shoe
	db := d.Context(ctx).Table(_ShoeTable).Where("shoe_id in ?", shoeIds)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "ShoesByIds  error.  shoeIds:(%+v);error:(%+v)", shoeIds, err)
		return nil, err
	}
	return respModel, nil
}

func (d *Dao) DeleteShoe(ctx context.Context, req *api.DeleteShoeReq) error {
	db := d.Context(ctx).Table(_ShoeTable)
	if len(req.Ids) != 0 {
		db.Where("shoe_id in ?", req.Ids)
	}
	db.Delete(&model.Shoe{})
	if err := db.Error; err != nil {
		log.Error(ctx, "DeleteShoe  error.  shoeId:(%+v);error:(%+v)", req, err)
		return err
	}
	return nil
}

func (d *Dao) SearchShoesPage(ctx context.Context, req *api.SearchShoesPageReq) ([]*model.Shoe, error) {
	var respModel []*model.Shoe
	db := d.Context(ctx).Table(_ShoeTable)
	if req.PageNum != 0 && req.PageSize != 0 {
		db.Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize))
	}
	BuildShoePageConditionDB(req.Condition, db)
	err := JsonBuild(req.ContentField, req.ContentFieldValues, db)
	if err != nil {
		return nil, err
	}
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchShoesPage  error.  req:(%+v);error:(%+v)", req, err)
		return nil, err
	}
	return respModel, nil
}

func (d *Dao) CountShoe(ctx context.Context, req *api.CountShoesReq) (int64, error) {
	var count int64
	db := d.Context(ctx).Table(_ShoeTable).Model(&model.Shoe{})
	BuildShoePageConditionDB(req.Condition, db)
	err := JsonBuild(req.ContentField, req.ContentFieldValues, db)
	if err != nil {
		return 0, err
	}
	db.Count(&count)
	if err := db.Error; err != nil {
		log.Error(ctx, "CountShoe  error.  req:(%+v);error:(%+v)", req, err)
		return 0, err
	}
	return count, nil
}
func BuildShoePageConditionDB(req *api.SearchShoeCondition, db *gorm.DB) {
	if req == nil {
		return
	}
	//if len(req.shoeIds) != 0 {
	//	db.Where("shoe_id in ?", req.shoeIds)
	//}
	if req.Material != "" {
		db.Where("material like ?", "%"+req.Material+"%")
	}
	if req.ShoeSize != "" {
		db.Where("shoe_size like ?", "%"+req.ShoeSize+"%")
	}
	if req.ShapeCode != "" {
		db.Where("shape_code like ?", "%"+req.ShapeCode+"%")
	}
	if req.IsHot != 0 {
		db.Where("is_hot = ?", req.IsHot)
	}
	if req.IsPresale != 0 {
		db.Where("is_presale = ?", req.IsPresale)
	}
}
func JsonBuild(contentField string, contentFieldValues []string, db *gorm.DB) error {
	//json字段查询
	if contentField != "" {
		zeroMap := commonTool.StructToDefaultValueMap(&api.Content{})
		v := zeroMap[contentField]
		var contentFieldInterfaceValues []interface{}
		switch v.(type) {
		case *uint32:
			for _, strV := range contentFieldValues {
				intV, err := strconv.ParseInt(strV, 32, 10)
				if err != nil {
					return err
				}
				contentFieldInterfaceValues = append(contentFieldInterfaceValues, uint32(intV))
			}
		case *string:
			for _, strV := range contentFieldValues {
				contentFieldInterfaceValues = append(contentFieldInterfaceValues, strV)
			}
		}
		queryContent := "content -> '$." + contentField + "' in ?"
		db.Where(queryContent, contentFieldInterfaceValues)
	}
	return nil
}
func (d *Dao) BatchAddShoe(ctx context.Context, Shoes []*model.Shoe) error {
	db := d.Context(ctx).Table(_ShoeTable).CreateInBatches(Shoes, 10)
	if err := db.Error; err != nil {
		log.Error(ctx, "BatchAddShoe  error.  Shoes:(%+v);error:(%+v)", Shoes, err)
		return err
	}
	return nil
}
