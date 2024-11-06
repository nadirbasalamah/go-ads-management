package ads

import "gorm.io/gorm"

type adsUseCase struct {
	adsRepository Repository
}

func NewAdsUseCase(repository Repository) UseCase {
	return &adsUseCase{
		adsRepository: repository,
	}
}

func (usecase *adsUseCase) GetAll() (*gorm.DB, error) {
	return usecase.adsRepository.GetAll()
}

func (usecase *adsUseCase) GetByID(id string) (Domain, error) {
	return usecase.adsRepository.GetByID(id)
}

func (usecase *adsUseCase) Create(adsReq *Domain) (Domain, error) {
	return usecase.adsRepository.Create(adsReq)
}

func (usecase *adsUseCase) Update(adsReq *Domain, id string) (Domain, error) {
	return usecase.adsRepository.Update(adsReq, id)
}

func (usecase *adsUseCase) Delete(id string) error {
	return usecase.adsRepository.Delete(id)
}

func (usecase *adsUseCase) Restore(id string) (Domain, error) {
	return usecase.adsRepository.Restore(id)
}

func (usecase *adsUseCase) ForceDelete(id string) error {
	return usecase.adsRepository.ForceDelete(id)
}
