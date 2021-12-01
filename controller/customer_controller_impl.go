package controller

// import (
// 	"kembanglah/helper"
// 	"kembanglah/model/web"
// 	"kembanglah/service"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type CustomerControllerImpl struct {
// 	CustomerService service.CustomerService
// }

// func NewCustomerController(customerService service.CustomerService) CustomerController {
// 	return &CustomerControllerImpl{
// 		CustomerService: customerService,
// 	}
// }

// func (u *CustomerControllerImpl) Login(ctx echo.Context) error {
// 	panic("implement me")
// }

// func (u *CustomerControllerImpl) Register(ctx echo.Context) error {
// 	customer := new(web.CustomerCreateRequest)
// 	if err := ctx.Bind(customer); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	if err := ctx.Validate(customer); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	customerResponse := u.CustomerService.Register(ctx.Request().Context(), *customer)

// 	return ctx.JSON(http.StatusOK, web.Response{
// 		Code:    http.StatusOK,
// 		Message: "succes create user",
// 		Data:    customerResponse,
// 	})
// }

// func (u *CustomerControllerImpl) FindByID(ctx echo.Context) error {
// 	customerId, err := strconv.Atoi(ctx.Param("customerId"))
// 	helper.PanicIfError(err)

// 	customerResponse := u.CustomerService.FindByID(ctx.Request().Context(), uint(customerId))
// 	return ctx.JSON(http.StatusOK, customerResponse)
// }

// func (u *CustomerControllerImpl) Update(ctx echo.Context) error {
// 	customer := new(web.CustomerUpdateRequest)

// 	customerId, err := strconv.Atoi(ctx.Param("customerId"))
// 	helper.PanicIfError(err)
// 	customer.ID = uint(customerId)

// 	if err := ctx.Bind(customer); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	if err := ctx.Validate(customer); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	userResponse := u.CustomerService.Update(ctx.Request().Context(), *customer)

// 	return ctx.JSON(http.StatusOK, web.Response{
// 		Code:    http.StatusOK,
// 		Message: "succes update user",
// 		Data:    userResponse,
// 	})
// }
