package repository

//
//import (
//	"20241223/model"
//	"be-golang-chapter-56/product-service/models"
//
//	"gorm.io/gorm"
//)
//
//type ProductRepository interface {
//	CreateProduct(product *model.Product) error
//	GetProductByID(id uint) (*model.Product, error)
//	GetAllProducts() ([]model.Product, error)
//}
//
//type productRepository struct {
//	db *gorm.DB
//}
//
//func NewProductRepository(db *gorm.DB) ProductRepository {
//	return &productRepository{db}
//}
//
//func (r *productRepository) CreateProduct(product *model.Product) error {
//	return r.db.Create(product).Error
//}
//
//func (r *productRepository) GetProductByID(id uint) (*model.Product, error) {
//	var product model.Product
//	err := r.db.First(&product, id).Error
//	return &product, err
//}
//
//func (r *productRepository) GetAllProducts() ([]model.Product, error) {
//	var products []model.Product
//	err := r.db.Find(&products).Error
//	return products, err
//}
