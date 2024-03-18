package request

import (
	"time"

	"github.com/cpartogi/menumanagement/internal/menu/model"
)

type MenuRequest struct {
	MenuName   string `json:"menu_name"`
	MenuDetail string `json:"menu_detail"`
	MenuPrice  int    `json:"menu_price"`
}

func (r *MenuRequest) FormatMenu(userId *int64) model.Menu {
	return model.Menu{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		MenuName:   r.MenuName,
		MenuDetail: r.MenuDetail,
		MenuPrice:  r.MenuPrice,
	}
}
