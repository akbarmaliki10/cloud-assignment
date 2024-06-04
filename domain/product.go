package domain

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float32 `json:"price" bson:"price"`
}

// here we create repository Contract
type ProductRepo interface {
	CreateProduct(*Product) error
	GetProduct(*uint) (*Product, error)
	GetAll() ([]*Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(*uint) error
}

type ProductUsecase interface {
	CreateProduct(*Product) error
	GetProduct(*string) (*Product, error)
	GetAll() ([]*Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(*string) error
}
