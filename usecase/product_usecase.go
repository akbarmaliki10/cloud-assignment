package usecase

import (
	"amitshekar-clean-arch/domain"
	"fmt"
	"strconv"
)

type ProductUsecaseImpl struct {
	repo domain.ProductRepo
}

// make a function that act like a constructor
func NewProductUsecase(ProductRepo domain.ProductRepo) domain.ProductUsecase {
	return &ProductUsecaseImpl{
		repo: ProductRepo,
	}
}

// receiver function or more like classes/struct method in python/java
func (u *ProductUsecaseImpl) CreateProduct(product *domain.Product) error {
	err := u.repo.CreateProduct(product)
	return err
}

func (u *ProductUsecaseImpl) GetProduct(id *string) (*domain.Product, error) {
	taskInt, err := strconv.Atoi(*id)
	if err != nil {
		fmt.Println("Error converting string to uint:", err)
		return nil, err
	}
	convertedId := uint(taskInt)
	product, err := u.repo.GetProduct(&convertedId)
	return product, err
}

func (u *ProductUsecaseImpl) GetAll() ([]*domain.Product, error) {
	products, err := u.repo.GetAll()
	return products, err
}

func (u *ProductUsecaseImpl) UpdateProduct(product *domain.Product) error {
	err := u.repo.UpdateProduct(product)
	return err
}

func (u *ProductUsecaseImpl) DeleteProduct(id *string) error {
	taskInt, err := strconv.Atoi(*id)
	if err != nil {
		fmt.Println("Error converting string to uint:", err)
		return err
	}
	convertedId := uint(taskInt)
	err = u.repo.DeleteProduct(&convertedId)
	return err
}
