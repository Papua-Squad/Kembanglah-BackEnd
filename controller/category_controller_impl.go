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
	if err := ctx.Bind(category); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := ctx.Validate(category); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	categoryResponse := controller.CategoryService.Create(ctx.Request().Context(), *category)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "success create product",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Update(ctx echo.Context) error {
	category := new(web.CategoryUpdateRequest)

	categoryID, err := strconv.Atoi(ctx.Param("categoryID"))
	helper.PanicIfError(err)
	category.ID = uint(categoryID)

	if err := ctx.Bind(category); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(category); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	categoryResponse := controller.CategoryService.Update(ctx.Request().Context(), *category)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "succes update category",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) FindByID(ctx echo.Context) error {
	categoryID, err := strconv.Atoi(ctx.Param("categoryID"))
	helper.PanicIfError(err)

	customerResponse := controller.CategoryService.FindByID(ctx.Request().Context(), uint(categoryID))
	return ctx.JSON(http.StatusOK, customerResponse)
}
