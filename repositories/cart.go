package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindToppingsById(ToppingID []int) ([]models.Topping, error)
	CartByTransaction(TrID int) ([]models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(Cart models.Cart) (models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	GetTransactionID() (models.Transaction, error)
	GetIDTransaction() (models.Transaction, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, ID).Error

	return cart, err
}

func (r *repository) FindToppingsById(ToppingID []int) ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Find(&toppings, ToppingID).Error

	return toppings, err
}

func (r *repository) CartByTransaction(TrID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Product").Preload("Topping").Find(&cart,"transaction_id = ?", TrID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) GetTransactionID() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Preload("Carts.Topping").Find(&transaction, "status = ?", "process").Error
	return transaction, err
}

func (r *repository) GetIDTransaction() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Preload("Carts.Topping").Find(&transaction, "status = ?", "process").Error
	return transaction, err
}