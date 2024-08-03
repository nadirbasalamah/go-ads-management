package services

import (
	"go-ads-management/models"
	"go-ads-management/repositories"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func InitCategoryService() CategoryService {
	return CategoryService{
		Repository: &repositories.CategoryRepositoryImpl{},
	}
}

func (cs *CategoryService) GetAll() ([]models.Category, error) {
	return cs.Repository.GetAll()
}

func (cs *CategoryService) GetByID(id string) (models.Category, error) {
	return cs.Repository.GetByID(id)
}

func (cs *CategoryService) Create(categoryInput models.CategoryInput) (models.Category, error) {
	return cs.Repository.Create(categoryInput)
}

func (cs *CategoryService) Update(categoryInput models.CategoryInput, id string) (models.Category, error) {
	return cs.Repository.Update(categoryInput, id)
}

func (cs *CategoryService) Delete(id string) error {
	return cs.Repository.Delete(id)
}
