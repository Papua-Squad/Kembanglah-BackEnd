package service

import (
	"context"
	"kembanglah/model/web"
)

type TransactionService interface {
	Save(ctx context.Context, request web.CreateTransactionRequest) (response web.TransactionResponse, err error)
	Update(ctx context.Context, request web.UpdateStatusTransactionRequest) (response web.TransactionResponse, err error)
	Delete(ctx context.Context, TransactionID uint) error
	FindBySeller(ctx context.Context, sellerID uint) (responses []web.TransactionResponse, err error)
	FindByCustomer(ctx context.Context, CustomerID uint) (responses []web.TransactionResponse, err error)
	FindByID(ctx context.Context, transactionID uint) (response web.TransactionResponse, err error)
	FindAll(ctx context.Context) (responses []web.TransactionResponse, err error)
}
