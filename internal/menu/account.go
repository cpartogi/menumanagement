package account

import (
	"context"
	"time"

	"github.com/cpartogi/izyai/internal/menu/model"
)

// Account denotes user's account object
type Account struct {
	ID           int64     `json:"id,omitempty"`
	Email        string    `json:"email,omitempty"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	PasswordHash string    `json:"password_hash,omitempty"`
	Status       Status    `json:"status,omitempty"`
	Name         string    `json:"name,omitempty"`
	ImageURL     string    `json:"image_url,omitempty"`
	CreateBy     int64     `json:"create_by,omitempty"`
	CreateTime   time.Time `json:"create_time,omitempty"`
	UpdateBy     int64     `json:"update_by,omitempty"`
	UpdateTime   time.Time `json:"update_time,omitempty"`
}

type Verification struct {
	ID               int64              `json:"id,omitempty"`
	AccountID        int64              `json:"account_id,omitempty"`
	Identity         string             `json:"identity,omitempty"`
	IdentityType     IdentityType       `json:"identity_type,omitempty"`
	VerificationCode string             `json:"verification_code,omitempty"`
	Status           VerificationStatus `json:"status,omitempty"`
	CreateBy         int64              `json:"create_by,omitempty"`
	CreateTime       time.Time          `json:"create_time,omitempty"`
	UpdateBy         int64              `json:"update_by,omitempty"`
	UpdateTime       time.Time          `json:"update_time,omitempty"`
}

// Registration denotes account registration
type Registration struct {
	Type    RegistrationType    `json:"type,omitempty"`
	Account AccountRegistration `json:"account,omitempty"`
	Company Company             `json:"company,omitempty"`
}

type AccountRegistration struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Company struct {
	Name   string `json:"name,omitempty"`
	Sector string `json:"sector,omitempty"`
}

// Credential denotes payload to authenticate an account
type Credential struct {
	Email       string
	PhoneNumber string
	Password    string
}

// Token denotes access token can be used as authentication method
type Token struct {
	Token string
}

type LoginResponse struct {
	Token      string `json:"token"`
	ExpiredAt  string `json:"expired_at"`
	CreatedAt  string `json:"created_at"`
	UserTypeID int64  `json:"user_type_id"`
	UserChatID int64  `json:"user_chat_id"`
}

type ListAccountFilter struct {
	Email       string
	PhoneNumber string
	Status      Status
}

type ListAccountVerificationFilter struct {
	Identity string
	Status   VerificationStatus
}

type ProvinceList struct {
	ProvinceData []ProvinceData
}

type ProvinceData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CityData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type DistrictData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserProfile struct {
	ID                 int64           `json:"id"`
	UserTypeID         int64           `json:"user_type_id"`
	PhoneNumber        string          `json:"phone_number"`
	Email              string          `json:"email"`
	Status             string          `json:"status"`
	Name               string          `json:"name"`
	ProfileImage       string          `json:"profile_image"`
	UserCompanyName    string          `json:"user_company_name"`
	UserCompanyPhone   string          `json:"user_company_phone"`
	UserCompanyAddress string          `json:"user_company_address"`
	IsPhoneVerified    bool            `json:"is_phone_verified"`
	IsEmailVerified    bool            `json:"is_email_verified"`
	Addresses          []UserAddresses `json:"addresses"`
	Company            UserCompany     `json:"company"`
}

type UserAddresses struct {
	ID                  int64  `json:"id"`
	ProvinceID          int64  `json:"province_id"`
	ProvinceName        string `json:"province_name"`
	CityID              int64  `json:"city_id"`
	CityName            string `json:"city_name"`
	DistrictID          int64  `json:"district_id"`
	DistrictName        string `json:"district_name"`
	PostalCode          int64  `json:"postal_code"`
	ReceiverName        string `json:"receiver_name"`
	ReceiverPhonenumber string `json:"receiver_phonenumber"`
	FullAddress         string `json:"full_address"`
	IsPrimary           bool   `json:"is_primary"`
}

type UserCompany struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	BusinessCategory string `json:"business_category"`
	Website          string `json:"website"`
	Linkedin         string `json:"linkedin"`
	Facebook         string `json:"facebook"`
	Twitter          string `json:"twitter"`
	Instagram        string `json:"instagram"`
}

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

// Status denotes account status
type Status int

const (
	StatusUnverified Status = 1
	StatusVerified   Status = 2
	StatusDeleted    Status = -1
)

func (s Status) Value() int {
	return int(s)
}

// RegistrationType denotes account registration type
type RegistrationType string

const (
	RegistrationTypePersonal RegistrationType = "personal"
	RegistrationTypeCompany  RegistrationType = "company"
)

type IdentityType int

const (
	IdentityTypeEmail       IdentityType = 1
	IdentityTypePhoneNumber IdentityType = 2
)

func (it IdentityType) Value() int {
	return int(it)
}

type VerificationStatus int

const (
	VerificationStatusUnverified VerificationStatus = 1
	VerificationStatusVerified   VerificationStatus = 2
	VerificationStatusRejected   VerificationStatus = 3
)

func (it VerificationStatus) Value() int {
	return int(it)
}

// Service denotes available method to access account module
type Service interface {
	GetMenus(ctx context.Context, filter map[string]string) ([]MenuData, error)
	GetMenuDetail(id *int64) ([]MenuDetailData, error)
	CreateMenu(mMenu *model.Menu) error
	UpdateMenu(id *int64, mMenu *model.Menu) error
	DeleteMenu(id *int64) error
}
