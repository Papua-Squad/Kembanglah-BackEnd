package controller

import (
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SellerControllerImpl struct {
	SellerService service.SellerService
}

func NewSellerController(sellerService service.SellerService) SellerController {
	return &SellerControllerImpl{
		SellerService: sellerService,
	}
}

func (u *SellerControllerImpl) Login(ctx echo.Context) error {
	panic("implement me")
}

func (u *SellerControllerImpl) Register(ctx echo.Context) error {

	user := new(web.SellerCreateRequest)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	userResponse := u.SellerService.Register(ctx.Request().Context(), *user)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "succes create user",
		Data:    userResponse,
	})
}

func (u *SellerControllerImpl) Update(ctx echo.Context) error {
	seller := new(web.SellerUpdateRequest)

	sellerId, err := strconv.Atoi(ctx.Param("sellerId"))
	helper.PanicIfError(err)
	seller.ID = uint(sellerId)

	if err := ctx.Bind(seller); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(seller); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	userResponse := u.SellerService.Update(ctx.Request().Context(), *seller)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "succes update user",
		Data:    userResponse,
	})
}

func (u *SellerControllerImpl) Delete(ctx echo.Context) error {
	sellerId, err := strconv.Atoi(ctx.Param("sellerId"))
	helper.PanicIfError(err)

	u.SellerService.Delete(ctx.Request().Context(), uint(sellerId))
	return ctx.JSON(http.StatusOK, "OK")
}

func (u *SellerControllerImpl) FindByID(ctx echo.Context) error {
	sellerId, err := strconv.Atoi(ctx.Param("sellerId"))
	helper.PanicIfError(err)

	sellerResponse := u.SellerService.FindByID(ctx.Request().Context(), uint(sellerId))
	return ctx.JSON(http.StatusOK, sellerResponse)
}

func (u *SellerControllerImpl) FindByUsername(ctx echo.Context) error {
	username := ctx.Param("username")

	sellerResponse := u.SellerService.FindByUsername(ctx.Request().Context(), string(username))
	return ctx.JSON(http.StatusOK, sellerResponse)
}

func (u *SellerControllerImpl) FindAll(ctx echo.Context) error {
	sellerResponse := u.SellerService.FindAll(ctx.Request().Context())
	return ctx.JSON(http.StatusOK, sellerResponse)
}
