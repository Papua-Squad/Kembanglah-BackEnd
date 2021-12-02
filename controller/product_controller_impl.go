package controller

import (
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"
	"strconv"

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

func (controller *ProductControllerImpl) Update(ctx echo.Context) error {
	product := new(web.ProductUpdateRequest)

	productID, err := strconv.Atoi(ctx.Param("productID"))
	helper.PanicIfError(err)
	product.ID = uint(productID)

	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	productResponse := controller.ProductService.Update(ctx.Request().Context(), *product)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "succes update product",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) Delete(ctx echo.Context) error {
	productID, err := strconv.Atoi(ctx.Param("productID"))
	helper.PanicIfError(err)

	controller.ProductService.Delete(ctx.Request().Context(), uint(productID))
	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "OK",
	})
}

func (controller *ProductControllerImpl) FindByID(ctx echo.Context) error {
	productID, err := strconv.Atoi(ctx.Param("productID"))
	helper.PanicIfError(err)

	customerResponse := controller.ProductService.FindByID(ctx.Request().Context(), uint(productID))
	return ctx.JSON(http.StatusOK, customerResponse)
}

func (controller *ProductControllerImpl) FindAll(ctx echo.Context) error {
	productResponse := controller.ProductService.FindAll(ctx.Request().Context())
	return ctx.JSON(http.StatusOK, productResponse)
}
