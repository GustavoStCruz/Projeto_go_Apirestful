package controller

import (
	"awesomeProject/model"
	"awesomeProject/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{productUseCase: usecase}

}
func (p *productController) Getproducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertProduct)

}

func (p *productController) GetproductsById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{Message: "id do produto não pode ser nulo"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{Message: "id do produto precisa ser um numero"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductId(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if product == nil {
		response := model.Response{Message: "Produto não foi encontrado na base de dados"}
		ctx.JSON(http.StatusNotFound, response)

	}

	ctx.JSON(http.StatusOK, product)
}
