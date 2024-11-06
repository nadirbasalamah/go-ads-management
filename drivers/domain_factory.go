package drivers

import (
	userDomain "go-ads-management/businesses/users"
	userDB "go-ads-management/drivers/mysql/users"

	"gorm.io/gorm"

	categoryDomain "go-ads-management/businesses/categories"
	categoryDB "go-ads-management/drivers/mysql/categories"

	adsDomain "go-ads-management/businesses/ads"
	adsDB "go-ads-management/drivers/mysql/ads"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewAdsRepository(conn *gorm.DB) adsDomain.Repository {
	return adsDB.NewMySQLRepository(conn)
}
