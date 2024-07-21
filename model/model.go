package model

//import "gorm.io/gorm"

// Kitten represents a kitten in the store.
type Kitten struct {
	//gorm.Model //create col(create_at,update_at,delete_at)
	Id      int `gorm:"type:int;primaryKey"`
	Name    string
	Age     int
	Color   string
	Breed   string
	Price   float64
	Count   uint
	InStock bool
	// Category Category
}

// func (Kitten) TableName() string {
//     return "kitten" // Specify the singular table name
// }

// // Category represents the category of a kitten.
// type Category struct {
// 	gorm.Model
// 	Name string `gorm:"unique;not null"`
// }

//Order represents an order made for kittens.
// type Order struct {
// 	//gorm.Model
// 	KittenID uint
// 	Kitten   Kitten
// 	Quantity int
// 	Total    float64
// }

// // User represents a user who can place orders.
type User struct {
	//gorm.Model
	Id       int    `gorm:"type:int;primaryKey"`
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	//Orders   []Order
}
