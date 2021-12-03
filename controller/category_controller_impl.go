package controller

import (
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(ctx echo.Context) error {

	category := new(web.CategoryRequest)
	if err := helper.BindAndValidate(ctx, category); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	categoryResponse, err := controller.CategoryService.Create(ctx.Request().Context(), *category)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Create Category",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Update(ctx echo.Context) error {
	category := new(web.CategoryUpdateRequest)

	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))
	category.ID = uint(categoryID)

	if err := helper.BindAndValidate(ctx, category); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	categoryResponse, err := controller.CategoryService.Update(ctx.Request().Context(), *category)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Update Category",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Delete(ctx echo.Context) error {
	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))

	if err := controller.CategoryService.Delete(ctx.Request().Context(), uint(categoryID)); err != nil {
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

func (controller *CategoryControllerImpl) FindByID(ctx echo.Context) error {
	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))

	categoryResponse, err := controller.CategoryService.FindByID(ctx.Request().Context(), uint(categoryID))
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
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) FindAll(ctx echo.Context) error {
	categoryResponse, err := controller.CategoryService.FindAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get all Categories Success",
		Data:    categoryResponse,
	})
}
