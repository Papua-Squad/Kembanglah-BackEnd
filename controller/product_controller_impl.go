package controller

import (
	"github.com/labstack/gommon/random"
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
	file, err := ctx.FormFile("image")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	pathFile, err := helper.SaveFile("product_"+random.Alphabetic, file)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	product := new(web.ProductRequest)
	product.ImageUrl = pathFile

	if err := helper.BindAndValidate(ctx, product); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	productResponse, err := controller.ProductService.Create(ctx.Request().Context(), *product)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Create Product",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) Update(ctx echo.Context) error {
	product := new(web.ProductUpdateRequest)

	productID, _ := strconv.Atoi(ctx.Param("productID"))
	product.ID = uint(productID)

	if err := helper.BindAndValidate(ctx, product); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	productResponse, err := controller.ProductService.Update(ctx.Request().Context(), *product)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Update Product",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) Delete(ctx echo.Context) error {
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	if err := controller.ProductService.Delete(ctx.Request().Context(), uint(productID)); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "OK",
	})
}

func (controller *ProductControllerImpl) FindByID(ctx echo.Context) error {
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	productResponse, err := controller.ProductService.FindByID(ctx.Request().Context(), uint(productID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Product By ID Success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) FindBySeller(ctx echo.Context) error {
	sellerID, _ := strconv.Atoi(ctx.Param("sellerID"))

	productResponse, err := controller.ProductService.FindByID(ctx.Request().Context(), uint(sellerID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Product By SellerID Success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) FindByCategory(ctx echo.Context) error {
	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))

	productResponse, err := controller.ProductService.FindByID(ctx.Request().Context(), uint(categoryID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Category By ID Success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) FindAll(ctx echo.Context) error {
	productResponse, err := controller.ProductService.FindAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get All Products Success",
		Data:    productResponse,
	})
}
