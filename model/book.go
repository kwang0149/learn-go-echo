package model

type Book struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"column:book_name" json:"book_name"`
	Description string `json:"description"`
}
