package models

import "time"

type User struct {
	ID           	uint 					`json:"id" gorm:"primaryKey"`
	FirstName 		string				`json:"firstName"`
	LastName 			string				`json:"lastName"`

	CreatedAt			time.Time			
	UpdatedAt			time.Time			
	DeletedAt			time.Time				
}	