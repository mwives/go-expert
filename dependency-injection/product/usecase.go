package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

// The DTO should be returned by the repository.
// The Product is being returned in order to keep the example simple
func (uc *ProductUseCase) GetProduct(id int) (Product, error) {
	return uc.repository.GetProduct(id)
}
