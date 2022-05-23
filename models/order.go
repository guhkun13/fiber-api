package models

import "time"

type Order struct {
	ID           		uint    				`json:"id" gorm:"primaryKey"`
	Name         		string  				`json:"name"`
	ProductRefer 		int     				`json:"productId"`
	Product      		Product 				`gorm:"foreignkey:ProductRefer"`
	UserRefer    		int     				`json:"userId"`
	User         		User    				`gorm:"foreignkey:UserRefer"`

	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			time.Time
}