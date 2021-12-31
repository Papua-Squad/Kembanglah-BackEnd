package controller

import (
	"github.com/labstack/echo/v4"
	"kembanglah/model/web"
	"kembanglah/service"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{TransactionService: transactionService}
}

func (controller *TransactionControllerImpl) Save(ctx echo.Context) error {
	transaction := new(web.CreateTransactionRequest)

}

func (controller *TransactionControllerImpl) Update(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) Delete(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) FindBySeller(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) FindByCustomer(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) FindByID(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) FindAll(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}
