package repository

import (
	"amitshekar-clean-arch/domain"
	"errors"

	"gorm.io/gorm"
)

type MysqlProductRepo struct {
	DB *gorm.DB
}

// New
// make a function that act like a constructor
func NewMysqlProductRepo(dB *gorm.DB) domain.ProductRepo {
	return &MysqlProductRepo{DB: dB}
}

// receiver function or more like classes/struct method in python/java
func (u *MysqlProductRepo) CreateProduct(product *domain.Product) error {
	if product.Name == "" {
		return errors.New("please fill product name")
	}
	err := u.DB.Create(product).Error
	return err
}

func (u *MysqlProductRepo) GetProduct(id *uint) (*domain.Product, error) {
	var product *domain.Product
	err := u.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, err
}

func (u *MysqlProductRepo) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
	err := u.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("document is empty")
	}
	return products, nil
}

// Update All with same name
func (u *MysqlProductRepo) UpdateProduct(product *domain.Product) error {
	err := u.DB.Model(&product).Where("id = ?", product.ID).Updates(product).Error

	if err != nil {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *MysqlProductRepo) DeleteProduct(name *uint) error {
	var product *domain.Product
	errFind := u.DB.Where("id = ?", name).First(&product).Error
	if errFind != nil {
		return errors.New("no matched document found for delete")
	}
	u.DB.Delete(&domain.Product{}, product.ID)
	return nil
}
