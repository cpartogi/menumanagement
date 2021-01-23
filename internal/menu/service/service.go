package service

import (
	"context"

	account "github.com/cpartogi/izyai/internal/menu"
	"github.com/cpartogi/izyai/internal/menu/model"
)

// Store represent the Account's repository contract
type Store interface {
	BeginTransaction() error
	Commit() error
	Rollback()
	GetAllMenus(ctx context.Context, filter map[string]string) ([]account.MenuData, error)
	GetDetailMenu(id *int64) ([]account.MenuDetailData, error)
	CreateMenu(m *model.Menu) error
	UpdateMenu(id *int64, m *model.Menu) error
	DeleteMenu(id *int64) error
}

type Service struct {
	store Store
}

// NewService register service for account domain
func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}
