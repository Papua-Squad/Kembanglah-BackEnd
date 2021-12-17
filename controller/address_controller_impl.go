package controller

import (
	"fmt"
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AddressControllerImpl struct {
	AddressService service.AddressService
}

func NewAddressController(addressService service.AddressService) AddressController {
	return &AddressControllerImpl{
		AddressService: addressService,
	}
}

func (controller *AddressControllerImpl) Create(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)
	userID := claims.Id
	fmt.Println(claims)

	address := new(web.AddressRequest)
	address.UserID = userID

	if err := helper.BindAndValidate(ctx, address); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	addressResponse, err := controller.AddressService.Create(ctx.Request().Context(), *address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Create Address",
		Data:    addressResponse,
	})
}

func (controller *AddressControllerImpl) Delete(ctx echo.Context) error {
	addressID, _ := strconv.Atoi(ctx.Param("addressID"))

	if err := controller.AddressService.Delete(ctx.Request().Context(), uint(addressID)); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Address Deleted",
	})
}

func (controller *AddressControllerImpl) FindByUser(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	addressResponse, err := controller.AddressService.FindByUser(ctx.Request().Context(), uint(userID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Address By User ID Success",
		Data:    addressResponse,
	})
}
