package account

import (
	"context"

	"github.com/cpartogi/menumanagement/internal/menu/model"
)

// Account denotes user's account object

type MenuData struct {
	ID        int64  `json:"id"`
	MenuName  string `json:"menu_name"`
	MenuPrice int64  `json:"menu_price"`
}

type MenuDetailData struct {
	ID         int64  `json:"id"`
	MenuName   string `json:"menu_name"`
	MenuDetail string `json:"menu_detail"`
	MenuPrice  int64  `json:"menu_price"`
}

// Service denotes available method to access account module
type Service interface {
	GetMenus(ctx context.Context, filter map[string]string) ([]MenuData, error)
	GetMenuDetail(id *int64) ([]MenuDetailData, error)
	CreateMenu(mMenu *model.Menu) error
	UpdateMenu(id *int64, mMenu *model.Menu) error
	DeleteMenu(id *int64) error
}
