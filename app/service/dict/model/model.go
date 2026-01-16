package model

import "gorm.io/gorm"

type Dict struct {
	gorm.Model
	DictId     string `json:"dict_id" gorm:"column:dict_id"`                  // 字典编号
	Name       string `json:"name" gorm:"column:name"`                        // 字典名称
	Identifier string `json:"identifier" gorm:"column:identifier"`            // 标识符
	DictType   string `json:"dict_type" gorm:"column:dict_type"`              // 字典类型
	DictValue  string `json:"dict_value" gorm:"column:dict_value;default:{}"` //字典值
	AppId      string `json:"app_id" gorm:"column:app_id"`
}

type UpdateDictReq struct {
	Name       string `json:"name"`       // 分类名称
	Identifier string `json:"identifier"` // 标识符
	DictType   string `json:"dict_type"`  // 字典类型
	DictValue  string `json:"dict_value"` //字典值
}
