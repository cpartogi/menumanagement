package service

import (
	"context"
	"log"

	account "github.com/cpartogi/menumanagement/internal/menu"
	"github.com/cpartogi/menumanagement/internal/menu/model"
)

func (s *Service) GetMenus(ctx context.Context, filter map[string]string) ([]account.MenuData, error) {
	return s.store.GetAllMenus(ctx, filter)
}

func (s *Service) GetMenuDetail(id *int64) ([]account.MenuDetailData, error) {
	return s.store.GetDetailMenu(id)
}

func (s *Service) CreateMenu(mMenu *model.Menu) error {
	var err error

	s.store.BeginTransaction()

	err = s.store.CreateMenu(mMenu)
	if err != nil {
		log.Println("Failed to store menu:", err)
		s.store.Rollback()
		return err
	}

	s.store.Commit()
	return err
}

func (s *Service) UpdateMenu(id *int64, mMenu *model.Menu) error {
	var err error

	s.store.BeginTransaction()

	err = s.store.UpdateMenu(id, mMenu)
	if err != nil {
		log.Println("Failed to update menu:", err)
		s.store.Rollback()
		return err
	}

	s.store.Commit()
	return err
}

func (s *Service) DeleteMenu(id *int64) error {
	var err error

	s.store.BeginTransaction()

	err = s.store.DeleteMenu(id)
	if err != nil {
		log.Println("Failed to delete menu:", err)
		s.store.Rollback()
		return err
	}

	s.store.Commit()
	return err
}
