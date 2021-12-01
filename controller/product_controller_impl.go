package controller

import (
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(ctx echo.Context) error {

	product := new(web.ProductRequest)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := ctx.Validate(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	productResponse := controller.ProductService.Create(ctx.Request().Context(), *product)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "success create product",
		Data:    productResponse,
	})
}
