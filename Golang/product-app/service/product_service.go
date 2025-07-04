package service

import (
	"github.com/pkg/errors"
	"product-app/domain"
	"product-app/repository"
	"product-app/service/dto"
)

type IProductService interface {
	Add(productCreateDto dto.ProductCreateDto) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) Add(productCreateDto dto.ProductCreateDto) error {
	validateErr := validateProductCreate(productCreateDto)
	if validateErr != nil {
		return validateErr
	}
	return productService.productRepository.AddProduct(domain.Product{
		Name:     productCreateDto.Name,
		Price:    productCreateDto.Price,
		Discount: productCreateDto.Discount,
		Store:    productCreateDto.Store,
	})
}

func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteById(productId)
}

func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetById(productId)
}

func (productService *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return productService.productRepository.UpdatePrice(productId, newPrice)
}

func (productService *ProductService) GetAllProducts() []domain.Product {
	return productService.productRepository.GetAllProducts()
}

func (productService *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(storeName)
}

func validateProductCreate(productCreateDto dto.ProductCreateDto) error {
	if productCreateDto.Discount > 70.0 {
		return errors.New("Discount can not be greater than 70")
	}

	return nil
}
