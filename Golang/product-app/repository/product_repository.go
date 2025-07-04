package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"product-app/domain"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")
	if err != nil {
		log.Error("Error while getting all products %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)

}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameSql := `SELECT * FROM products WHERE store = $1`

	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)
	if err != nil {
		log.Error("Error while getting all products %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)

}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insert_sql := `INSERT INTO products(name, price,discount,store) VALUES($1, $2, $3, $4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Error("Failed to add new product %v", err)
		return err
	}

	log.Info(fmt.Printf("Added new product %v\n", addNewProduct))
	return nil
}

func (ProductRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()

	getByIdSql := `SELECT * FROM products WHERE id = $1`

	productRow := ProductRepository.dbPool.QueryRow(ctx, getByIdSql, productId)

	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanErr := productRow.Scan(&id, &name, &price, &discount, &store)
	if scanErr != nil {
		log.Errorf("Failed to scan product with id %d: %v", productId, scanErr)
		return domain.Product{}, fmt.Errorf("product with id %d not found", productId)
	}
	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil

}

func (ProductRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()

	_, getErr := ProductRepository.GetById(productId)
	if getErr != nil {
		return errors.New("Product not found !")
	}

	deleteByIdSql := `DELETE FROM products WHERE id = $1`

	_, err := ProductRepository.dbPool.Exec(ctx, deleteByIdSql, productId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while deleting product with id %d", productId))
	}

	log.Info("Product deleted !")
	return nil
}

func (ProductRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()

	updateSql := `UPDATE products SET price = $1 WHERE id = $2`

	_, err := ProductRepository.dbPool.Exec(ctx, updateSql, newPrice, productId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while updating product with id %d", productId))
	}

	log.Info("Product %d price updated with new price %v !", productId, newPrice)
	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {

	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}

	return products
}
