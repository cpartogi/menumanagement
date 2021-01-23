package model

import "time"

type Province struct {
	ID         int64      `gorm:"column:id"`
	Name       string     `gorm:"column:name"`
	CreateBy   int64      `gorm:"column:create_by"`
	CreateTime time.Time  `gorm:"column:create_time"`
	UpdateBy   int64      `gorm:"column:update_by"`
	UpdateTime time.Time  `gorm:"column:update_time"`
	DeleteBy   *int64     `gorm:"column:delete_by"`
	DeleteTime *time.Time `gorm:"column:delete_time"`
}
