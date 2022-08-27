package models

import "time"

type Product struct {
	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title      string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price      int                  `json:"price" form:"price" gorm:"type: int"`
	Image      string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	CreatedAt  time.Time            `json:"-"`
	UpdatedAt  time.Time            `json:"-"`
}

type ProductResponse struct {
	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title      string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price      int                  `json:"price" form:"price" gorm:"type: int"`
	Image      string               `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type ProductCartResponse struct {
	ID         int                  `json:"id"`
	Title      string               `json:"title"`
	Price      int                  `json:"price"`
	Image      string               `json:"image"`
}

func (ProductResponse) TableName() string {
	return "products"
}


// func (ProductCartResponse) TableName() string {
// 	return "products"
// }