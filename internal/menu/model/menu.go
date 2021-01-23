package model

import "time"

type Menu struct {
	id         int64     `gorm:"column:id"`
	MenuName   string    `gorm:"column:menu_name"`
	MenuDetail string    `gorm:"column:menu_detail"`
	MenuPrice  int       `gorm:"column:menu_price"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
