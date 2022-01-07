package controller

import (
	"github.com/labstack/echo/v4"
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"
	"strconv"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{TransactionService: transactionService}
}

func (controller *TransactionControllerImpl) Save(ctx echo.Context) error {
	transaction := new(web.CreateTransactionRequest)

	if err := helper.BindAndValidate(ctx, transaction); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	transactionResponse, err := controller.TransactionService.Save(ctx.Request().Context(), *transaction)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, web.Response{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Save Transaction",
		Data:    transactionResponse,
	})
}

func (controller *TransactionControllerImpl) Update(ctx echo.Context) error {
	transaction := new(web.UpdateStatusTransactionRequest)
	transactionID, _ := strconv.Atoi(ctx.Param("transactionID"))
	transaction.ID = uint(transactionID)

	if err := helper.BindAndValidate(ctx, transaction); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	transactionResponse, err := controller.TransactionService.Update(ctx.Request().Context(), *transaction)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, web.Response{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Update Transaction",
		Data:    transactionResponse,
	})
}

func (controller *TransactionControllerImpl) Delete(ctx echo.Context) error {
	transactionID, _ := strconv.Atoi(ctx.Param("transactionID"))

	if err := controller.TransactionService.Delete(ctx.Request().Context(), uint(transactionID)); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success delete Transaction",
	})
}

func (controller *TransactionControllerImpl) FindBySeller(ctx echo.Context) error {
	sellerID, _ := strconv.Atoi(ctx.Param("sellerID"))

	transactionResponse, err := controller.TransactionService.FindBySeller(ctx.Request().Context(), uint(sellerID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Transaction By sellerID Success",
		Data:    transactionResponse,
	})
}

func (controller *TransactionControllerImpl) FindByCustomer(ctx echo.Context) error {
	customerID, _ := strconv.Atoi(ctx.Param("customerID"))

	transactionResponse, err := controller.TransactionService.FindByCustomer(ctx.Request().Context(), uint(customerID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Transaction By customerId Success",
		Data:    transactionResponse,
	})
}

func (controller *TransactionControllerImpl) FindByID(ctx echo.Context) error {
	transactionID, _ := strconv.Atoi(ctx.Param("transactionID"))

	transactionResponse, err := controller.TransactionService.FindByCustomer(ctx.Request().Context(), uint(transactionID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get Transaction By transactionID Success",
		Data:    transactionResponse,
	})
}

func (controller *TransactionControllerImpl) FindAll(ctx echo.Context) error {
	transactionResponse, err := controller.TransactionService.FindAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get All Transaction Success",
		Data:    transactionResponse,
	})
}
