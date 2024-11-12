package categories

import "context"

type categoryUseCase struct {
	categoryRepository Repository
}

func NewCategoryUseCase(repository Repository) UseCase {
	return &categoryUseCase{
		categoryRepository: repository,
	}
}

func (usecase *categoryUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.categoryRepository.GetAll(ctx)
}

func (usecase *categoryUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	return usecase.categoryRepository.GetByID(ctx, id)
}

func (usecase *categoryUseCase) Create(ctx context.Context, categoryReq *Domain) (Domain, error) {
	return usecase.categoryRepository.Create(ctx, categoryReq)
}

func (usecase *categoryUseCase) Update(ctx context.Context, categoryReq *Domain, id int) (Domain, error) {
	return usecase.categoryRepository.Update(ctx, categoryReq, id)
}

func (usecase *categoryUseCase) Delete(ctx context.Context, id int) error {
	return usecase.categoryRepository.Delete(ctx, id)
}
