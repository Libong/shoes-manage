package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryId string `json:"category_id" gorm:"column:category_id"` // 分类编号
	Name       string `json:"name" gorm:"column:name"`               // 分类名称
	Identifier string `json:"identifier" gorm:"column:identifier"`   // 标识符
	AppId      string `json:"app_id" gorm:"column:app_id"`
}

type UpdateCategoryReq struct {
	Name       string `json:"name"`       // 分类名称
	Identifier string `json:"identifier"` // 标识符
}
