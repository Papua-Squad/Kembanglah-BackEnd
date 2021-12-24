package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt"
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

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)
	seller := claims.Id
	fmt.Println(claims)

	pathFile, err := helper.SaveFile("product_"+random.String(8, random.Alphabetic)+".jpg", file)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	product := new(web.ProductRequest)
	product.ImageUrl = "http://159.223.82.24:3000/files/" + pathFile
	product.SellerID = seller

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
		Message: "Product Deleted",
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

	productResponse, err := controller.ProductService.FindBySeller(ctx.Request().Context(), uint(sellerID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Product By Seller ID Success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) FindByCategory(ctx echo.Context) error {
	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))

	productResponse, err := controller.ProductService.FindByCategory(ctx.Request().Context(), uint(categoryID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Product By Category ID Success",
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
