package categories

import (
	"context"
	"go-ads-management/businesses/categories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) categories.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	var records []Category

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []categories.Domain{}

	for _, category := range records {
		categories = append(categories, category.ToDomain())
	}

	return categories, nil
}

func (cr *categoryRepository) GetByID(ctx context.Context, id string) (categories.Domain, error) {
	var category Category

	if err := cr.conn.WithContext(ctx).First(&category, "id = ?", id).Error; err != nil {
		return categories.Domain{}, err
	}

	return category.ToDomain(), nil
}

func (cr *categoryRepository) Create(ctx context.Context, categoryReq *categories.Domain) (categories.Domain, error) {
	record := FromDomain(categoryReq)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return categories.Domain{}, err
	}

	if err := result.WithContext(ctx).Last(&record).Error; err != nil {
		return categories.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *categoryRepository) Update(ctx context.Context, categoryReq *categories.Domain, id string) (categories.Domain, error) {
	category, err := cr.GetByID(ctx, id)

	if err != nil {
		return categories.Domain{}, err
	}

	updatedCategory := FromDomain(&category)

	updatedCategory.Name = categoryReq.Name

	if err := cr.conn.WithContext(ctx).Save(&updatedCategory).Error; err != nil {
		return categories.Domain{}, err
	}

	return updatedCategory.ToDomain(), nil
}

func (cr *categoryRepository) Delete(ctx context.Context, id string) error {
	category, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedCategory := FromDomain(&category)

	if err := cr.conn.WithContext(ctx).Delete(&deletedCategory).Error; err != nil {
		return err
	}

	return nil
}
