package models

import "time"

type Product struct {
	ID           			uint   				`json:"id" gorm:"primaryKey"`
	Name         			string 				`json:"name"`
	SerialNumber 			string 				`json:"serialNumber"`

	CreatedAt 				time.Time
	UpdatedAt 				time.Time
	DeletedAt 				time.Time
}
