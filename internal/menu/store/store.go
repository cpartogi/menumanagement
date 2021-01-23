package repository

import (
	"context"

	account "github.com/cpartogi/izyai/internal/menu"
	"github.com/cpartogi/izyai/internal/menu/model"

	"github.com/jinzhu/gorm"
)

type Store struct {
	Conn *gorm.DB
	DB   *gorm.DB
}

func (s *Store) BeginTransaction() error {
	s.Conn = s.Conn.Begin()
	return s.Conn.Error
}

func (s *Store) Commit() error {
	s.Conn.Commit()
	s.Conn = s.DB
	return s.Conn.Error
}

func (s *Store) Rollback() {
	s.Conn.Rollback()
	s.Conn = s.DB
}

// NewStore will create an object that represent the Repository interface
func NewStore(dbConn *gorm.DB) *Store {
	return &Store{Conn: dbConn, DB: dbConn}
}

func (s *Store) GetAllMenus(ctx context.Context, filter map[string]string) ([]account.MenuData, error) {
	var result []account.MenuData

	query := s.Conn.Table("menus")
	query = query.Select("id, menu_name, menu_price")
	if filter["name"] != "" {
		query = query.Where("menu_name ILIKE '%" + filter["name"] + "%'")
	}
	query.Find(&result)

	if query.RecordNotFound() {
		return result, account.ErrMenuNotFound
	}

	return result, query.Error
}

func (s *Store) GetDetailMenu(id *int64) ([]account.MenuDetailData, error) {
	var result []account.MenuDetailData

	query := s.Conn.Table("menus")
	query = query.Select("id, menu_name, menu_detail, menu_price")
	query = query.Where("id = ?", *id)

	query.Find(&result)
	if query.RecordNotFound() {
		return result, account.ErrMenuNotFound
	}

	return result, query.Error
}

func (s *Store) CreateMenu(m *model.Menu) error {
	query := s.Conn.Create(m)

	return query.Error
}

func (s *Store) UpdateMenu(id *int64, m *model.Menu) error {
	query := s.Conn.Model(m).Where("id = ?", *id).Updates(*m)

	return query.Error
}

func (s *Store) DeleteMenu(id *int64) error {

	query := s.Conn.Table("menus")
	query = query.Delete("id, menu_name, menu_detail, menu_price")
	query = query.Where("id = ?", *id)

	return query.Error
}
