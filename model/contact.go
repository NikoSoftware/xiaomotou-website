package model

import (
	"time"
)

type Contact struct {
	Id         int       `gorm:"column:id" `
	Name       string    `gorm:"column:name" form:"name" validate:"required" reg_error_info:"名称不能为空"`
	Phone      string    `gorm:"column:phone" form:"phone" validate:"required,PhoneValidationErrors" reg_error_info:"手机号异常"`
	Title      string    `gorm:"column:title" form:"title" validate:"required" reg_error_info:"标题不能为空"`
	Content    string    `gorm:"column:content" form:"content" validate:"required" reg_error_info:"内容不能为空"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Contact) TableName() string {
	return "contact"
}
