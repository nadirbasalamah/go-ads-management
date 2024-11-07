package ads

import (
	"context"

	"gorm.io/gorm"
)

type adsUseCase struct {
	adsRepository Repository
}

func NewAdsUseCase(repository Repository) UseCase {
	return &adsUseCase{
		adsRepository: repository,
	}
}

func (usecase *adsUseCase) GetAll(ctx context.Context) (*gorm.DB, error) {
	return usecase.adsRepository.GetAll(ctx)
}

func (usecase *adsUseCase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.adsRepository.GetByID(ctx, id)
}

func (usecase *adsUseCase) GetByCategory(ctx context.Context, categoryID string) (*gorm.DB, error) {
	return usecase.adsRepository.GetByCategory(ctx, categoryID)
}

func (usecase *adsUseCase) GetByUser(ctx context.Context) (*gorm.DB, error) {
	return usecase.adsRepository.GetByUser(ctx)
}

func (usecase *adsUseCase) GetTrashed(ctx context.Context) (*gorm.DB, error) {
	return usecase.adsRepository.GetTrashed(ctx)
}

func (usecase *adsUseCase) Create(ctx context.Context, adsReq *Domain) (Domain, error) {
	return usecase.adsRepository.Create(ctx, adsReq)
}

func (usecase *adsUseCase) Update(ctx context.Context, adsReq *Domain, id string) (Domain, error) {
	return usecase.adsRepository.Update(ctx, adsReq, id)
}

func (usecase *adsUseCase) Delete(ctx context.Context, id string) error {
	return usecase.adsRepository.Delete(ctx, id)
}

func (usecase *adsUseCase) Restore(ctx context.Context, id string) (Domain, error) {
	return usecase.adsRepository.Restore(ctx, id)
}

func (usecase *adsUseCase) ForceDelete(ctx context.Context, id string) error {
	return usecase.adsRepository.ForceDelete(ctx, id)
}
