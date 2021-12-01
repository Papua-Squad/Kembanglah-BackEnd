package controller

// import (
// 	"kembanglah/helper"
// 	"kembanglah/model/web"
// 	"kembanglah/service"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type SellerControllerImpl struct {
// 	SellerService service.UserService
// }

// func NewSellerController(sellerService service.UserService) SellerController {
// 	return &SellerControllerImpl{
// 		SellerService: sellerService,
// 	}
// }

// func (controller *SellerControllerImpl) Login(ctx echo.Context) error {
// 	panic("implement me")
// }

// func (controller *SellerControllerImpl) Register(ctx echo.Context) error {

// 	seller := new(web.SellerRegisterRequest)
// 	if err := ctx.Bind(seller); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, web.Response{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 	}

// 	if err := ctx.Validate(seller); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, web.Response{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 	}

// 	sellerResponse := controller.SellerService.Register(ctx.Request().Context(), *seller)

// 	return ctx.JSON(http.StatusOK, web.Response{
// 		Code:    http.StatusOK,
// 		Message: "success create seller",
// 		Data:    sellerResponse,
// 	})
// }

// func (controller *SellerControllerImpl) Update(ctx echo.Context) error {
// 	seller := new(web.SellerUpdateRequest)

// 	sellerId, err := strconv.Atoi(ctx.Param("sellerId"))
// 	helper.PanicIfError(err)
// 	seller.ID = uint(sellerId)

// 	if err := ctx.Bind(seller); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	if err := ctx.Validate(seller); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	userResponse := controller.SellerService.Update(ctx.Request().Context(), *seller)

// 	return ctx.JSON(http.StatusOK, web.Response{
// 		Code:    http.StatusOK,
// 		Message: "succes update user",
// 		Data:    userResponse,
// 	})
// }

// func (controller *SellerControllerImpl) Delete(ctx echo.Context) error {
// 	sellerId, err := strconv.Atoi(ctx.Param("sellerId"))
// 	helper.PanicIfError(err)

// 	controller.SellerService.Delete(ctx.Request().Context(), uint(sellerId))
// 	return ctx.JSON(http.StatusOK, "OK")
// }

// func (controller *SellerControllerImpl) FindByID(ctx echo.Context) error {
// 	sellerId, err := strconv.Atoi(ctx.Param("sellerId"))
// 	helper.PanicIfError(err)

// 	sellerResponse := controller.SellerService.FindByID(ctx.Request().Context(), uint(sellerId))
// 	return ctx.JSON(http.StatusOK, sellerResponse)
// }

// func (controller *SellerControllerImpl) FindByUsername(ctx echo.Context) error {
// 	username := ctx.Param("username")

// 	sellerResponse := controller.SellerService.FindByUsername(ctx.Request().Context(), string(username))
// 	return ctx.JSON(http.StatusOK, sellerResponse)
// }

// func (controller *SellerControllerImpl) FindAll(ctx echo.Context) error {
// 	sellerResponse := controller.SellerService.FindAll(ctx.Request().Context())
// 	return ctx.JSON(http.StatusOK, sellerResponse)
// }
