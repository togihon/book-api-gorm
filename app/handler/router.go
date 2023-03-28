package handler

import (
	"book-gorm/app/services"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	//get book by id
	router.GET("/books/", services.GetSemuaBuku)                //get all books
	router.GET("/books/:bookID", services.GetBukuBerdasarkanID) //get book by id
	router.POST("/books/", services.TambahBuku)                 //create book
	router.PUT("/books/:bookID", services.EditBuku)             //update book by id
	router.DELETE("/books/:bookID", services.HapusBuku)         //delete book by id
	return router
}
