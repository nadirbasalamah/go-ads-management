package categories

type categoryUseCase struct {
	categoryRepository Repository
}

func NewCategoryUseCase(repository Repository) UseCase {
	return &categoryUseCase{
		categoryRepository: repository,
	}
}

func (usecase *categoryUseCase) GetAll() ([]Domain, error) {
	return usecase.categoryRepository.GetAll()
}

func (usecase *categoryUseCase) GetByID(id string) (Domain, error) {
	return usecase.categoryRepository.GetByID(id)
}

func (usecase *categoryUseCase) Create(categoryReq *Domain) (Domain, error) {
	return usecase.categoryRepository.Create(categoryReq)
}

func (usecase *categoryUseCase) Update(categoryReq *Domain, id string) (Domain, error) {
	return usecase.categoryRepository.Update(categoryReq, id)
}

func (usecase *categoryUseCase) Delete(id string) error {
	return usecase.categoryRepository.Delete(id)
}
