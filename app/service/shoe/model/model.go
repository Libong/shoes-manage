package model

import "gorm.io/gorm"

/*
	数据库
*/

type Shoe struct {
	gorm.Model
	ShoeId    string `gorm:"column:shoe_id" json:"shoe_id"`                 //鞋子编号
	ShapeCode string `gorm:"column:shape_code" json:"shape_code"`           //形体编码
	Material  string `gorm:"column:material" json:"material"`               //材质
	ShoeSize  string `gorm:"column:shoe_size" json:"shoe_size"`             //鞋子码数
	IsHot     uint32 `gorm:"column:is_hot;default:2" json:"is_hot"`         //是否热门 1是 2否
	IsPresale uint32 `gorm:"column:is_presale;default:2" json:"is_presale"` //是否预售 1是 2否
	Pictures  string `gorm:"column:pictures" json:"pictures"`               // 图片
	Videos    string `gorm:"column:videos" json:"videos"`                   // 视频
}
type AccountShoe struct {
	gorm.Model
	ShoeId    string `gorm:"column:shoe_id" json:"shoe_id"` //鞋子编号
	AccountId string `gorm:"column:account_id" json:"account_id"`
}
type UpdateShoeReq struct {
	ShapeCode string `gorm:"column:shape_code" json:"shape_code"` //角色名称
	Material  string `gorm:"column:material" json:"material"`     //描述
	ShoeSize  string `gorm:"column:shoe_size" json:"shoe_size"`   //可分配角色
	IsHot     uint32 `gorm:"column:is_hot" json:"is_hot"`         //是否热门 1是 2否
	IsPresale uint32 `gorm:"column:is_presale" json:"is_presale"` //是否预售 1是 2否
	Pictures  string `gorm:"column:pictures" json:"pictures"`
	Videos    string `gorm:"column:videos" json:"videos"`
}
