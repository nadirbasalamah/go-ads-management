package users

import (
	"context"
	"errors"
	"go-ads-management/app/middlewares"
	"go-ads-management/businesses/users"
	"go-ads-management/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(ctx context.Context, userReq *users.Domain) (users.Domain, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return users.Domain{}, err
	}

	record := FromDomain(userReq)

	record.Password = string(password)

	result := ur.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	if err := result.WithContext(ctx).Last(&record).Error; err != nil {
		return users.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ur *userRepository) CreateAdmin(ctx context.Context, userReq *users.Domain) (users.Domain, error) {
	var count int64

	err := ur.conn.WithContext(ctx).Model(&User{}).Where("role = ?", utils.ROLE_ADMIN).Count(&count).Error

	if err != nil {
		return users.Domain{}, err
	}

	if count == 1 {
		return users.Domain{}, errors.New("admin already exists")
	}

	user, err := ur.Register(ctx, userReq)

	if err != nil {
		return users.Domain{}, err
	}

	return user, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, userReq *users.Domain) (users.Domain, error) {
	var user User

	err := ur.conn.WithContext(ctx).First(&user, "email = ?", userReq.Email).Error

	if err != nil {
		return users.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (ur *userRepository) GetUserInfo(ctx context.Context) (users.Domain, error) {
	id, err := middlewares.GetUserID(ctx)

	if err != nil {
		return users.Domain{}, errors.New("invalid token")
	}

	var user User

	err = ur.conn.WithContext(ctx).First(&user, "id = ?", id).Error

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}
