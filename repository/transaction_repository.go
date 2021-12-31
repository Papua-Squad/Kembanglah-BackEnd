package repository

import (
	"context"
	domain "kembanglah/model/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	Update(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	Delete(ctx context.Context, transaction domain.Transaction) error
	FindBySeller(ctx context.Context, sellerID uint) (transactions []domain.Transaction, err error)
	FindByCustomer(ctx context.Context, customerID uint) (transactions []domain.Transaction, err error)
	FindByID(ctx context.Context, transactionID uint) (transaction domain.Transaction, err error)
	FindAll(ctx context.Context) (transactions []domain.Transaction, err error)
}
