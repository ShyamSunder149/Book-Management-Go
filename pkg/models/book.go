package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shyamsunder149/go-bookstore/pkg/config"
)

var db *gorm.DB 

type Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
	author string `json:"author"`
	publication string  `json:"publication"`\
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db := db.Where("ID=?",Id).Delete(book)
	return book
}

