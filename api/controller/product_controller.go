package controller

import (
	"amitshekar-clean-arch/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product domain.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := pc.ProductUsecase.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *ProductController) GetProduct(ctx *gin.Context) {
	var taskId string = ctx.Param("id")
	product, err := pc.ProductUsecase.GetProduct(&taskId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAll(ctx *gin.Context) {
	products, err := pc.ProductUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	taskId := ctx.Param("id")
	taskInt, err := strconv.Atoi(taskId)
	if err != nil {
		fmt.Println("Error converting string to uint:", err)
		return
	}
	var product domain.Product
	product.ID = uint(taskInt)
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = pc.ProductUsecase.UpdateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	taskId := ctx.Param("id")
	err := pc.ProductUsecase.DeleteProduct(&taskId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
