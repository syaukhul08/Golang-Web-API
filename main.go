package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	//CRUD
	// book := book.Book{}
	// book.Title = "Biologi"
	// book.Price = 90000
	// book.Discount = 20
	// book.Rating = 5
	// book.Description = "Buku pelajaran biologi"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("===================")
	// }

	var book book.Book

	err = db.Debug().Where("id = ?", 4).Find(&book).Error
	if err != nil {
		fmt.Println("==================")
		fmt.Println("Error finding book record")
		fmt.Println("===================")
	}

	book.Title = "Sejarah Kemerdekaan Indonesia"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("==================")
		fmt.Println("Error updating book record")
		fmt.Println("===================")
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBookHandler)

	router.Run(":8008")
}
