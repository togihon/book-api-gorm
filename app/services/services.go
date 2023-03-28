package services

import (
	"book-gorm/app/entity"
	"book-gorm/app/repo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSemuaBuku(ctx *gin.Context) {
	result, err := repo.GetSemuaData()

	if err != nil {
		panic(err)
	}

	if len(result) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Tidak ada data buku tersimpan",
		})
	}

	if len(result) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    result,
		})
	}
}

func GetBukuBerdasarkanID(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))

	result, err := repo.GetDataBerdasarkanID(bookID)

	if err != nil {
		panic(err)
	}

	if len(result) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": fmt.Sprintf("Buku dengan id %d tidak ditemukan", bookID),
		})
	}

	if len(result) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    result,
		})
	}
}

func TambahBuku(ctx *gin.Context) {
	var requestBody entity.BookGorm

	if err := ctx.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	result, err := repo.CreateBuku(requestBody.Title, requestBody.Author)
	if err != nil {
		panic(err)
	}

	if len(result) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Data buku gagal di-input",
		})
	}

	if len(result) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    result,
		})
	}

}

func EditBuku(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	var requestBody entity.BookGorm

	if err := ctx.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	result, err := repo.UpdateBuku(bookID, requestBody.Title, requestBody.Author)
	if err != nil {
		panic(err)
	}

	if len(result) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Data buku gagal di-update",
		})
	}

	if len(result) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    result,
		})
	}
}

func HapusBuku(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	result, err := repo.DeleteBuku(bookID)
	if err != nil {
		panic(err)
	}

	if result == "Gagal" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Data buku gagal dihapus",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    fmt.Sprintf("Data buku dengan id %d berhasil dihapus", bookID),
	})
}
