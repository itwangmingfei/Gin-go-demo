package models

import "github.com/jinzhu/gorm"

type Content struct {
	gorm.Model
	Uid uint64 `gorm:"not null;commit:'用户ID'"`
	Title string `gorm:"size:50;not null;commit:'标题'"`
	Contents string `gorm:"type:text;commit:'内容'"`
}
func(c Content) TableName()string{
	return "content"
}
