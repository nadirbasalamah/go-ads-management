package users

import (
	"context"
	"go-ads-management/utils"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	CompanyName string
	Address     string
	Username    string
	Email       string
	Password    string
	Role        utils.Role
}

type UseCase interface {
	Register(ctx context.Context, userReq *Domain) (Domain, error)
	Login(ctx context.Context, userReq *Domain) (string, error)
	GetUserInfo(ctx context.Context) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, userReq *Domain) (Domain, error)
	CreateAdmin(ctx context.Context, userReq *Domain) (Domain, error)
	GetByEmail(ctx context.Context, userReq *Domain) (Domain, error)
	GetUserInfo(ctx context.Context) (Domain, error)
}
