package model

import (
	"time"
)

type Contact struct {
	Id         int       `gorm:"column:id" `
	Name       string    `gorm:"column:name" form:"name" validate:"required"`
	Phone      string    `gorm:"column:phone" form:"phone" validate:"required,PhoneValidationErrors"`
	Title      string    `gorm:"column:title" form:"title" validate:"required"`
	Content    string    `gorm:"column:content" form:"content" validate:"required"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Contact) TableName() string {
	return "contact"
}
