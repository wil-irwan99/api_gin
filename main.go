package main

import (
	"fmt"
	"gin_latihan/model"
	"gin_latihan/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	pdtRepo := repo.NewProductRepository()
	ctgRepo := repo.NewCategoryRepository()

	r := gin.Default()

	r.POST("/product", func(c *gin.Context) {
		var result []model.Product
		var newProduct model.Product
		if err := c.BindJSON(&newProduct); err != nil { //query harus json
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    newProduct,
			})
		}
		result = pdtRepo.Add(&newProduct)
		fmt.Println(result)
	})

	r.POST("/category", func(c *gin.Context) {
		var result []model.Category
		var newCategory model.Category
		if err := c.BindJSON(&newCategory); err != nil { //query harus json
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    newCategory,
			})
		}
		result = ctgRepo.Add(&newCategory)
		fmt.Println(result)
	})

	r.GET("/category", func(c *gin.Context) {
		var category model.Category
		if err := c.BindJSON(&category); err != nil { //query harus json
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    category,
			})
		}
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
