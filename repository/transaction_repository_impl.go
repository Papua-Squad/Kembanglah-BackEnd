package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type TransactionRepositoryImpl struct {
	server *app.Server
}

func NewTransactionRepositoryImpl(server *app.Server) TransactionRepository {
	return &TransactionRepositoryImpl{server: server}
}

func (repository *TransactionRepositoryImpl) Save(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	return transaction, repository.server.DB.WithContext(ctx).Create(&transaction).Error
}

func (repository *TransactionRepositoryImpl) Update(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	return transaction, repository.server.DB.WithContext(ctx).Model(&transaction).Updates(&transaction).Error
}

func (repository *TransactionRepositoryImpl) Delete(ctx context.Context, transaction domain.Transaction) error {
	return repository.server.DB.WithContext(ctx).Delete(&transaction).Error
}

func (repository *TransactionRepositoryImpl) FindBySeller(ctx context.Context, sellerID uint) (transactions []domain.Transaction, err error) {
	return transactions, repository.server.DB.WithContext(ctx).Where("seller_id = ?", sellerID).Find(&transactions).Error
}

func (repository *TransactionRepositoryImpl) FindByCustomer(ctx context.Context, customerID uint) (transactions []domain.Transaction, err error) {
	return transactions, repository.server.DB.WithContext(ctx).Where("customer_id = ?", customerID).Find(&transactions).Error
}

func (repository *TransactionRepositoryImpl) FindByID(ctx context.Context, transactionID uint) (transaction domain.Transaction, err error) {
	return transaction, repository.server.DB.WithContext(ctx).First(&transaction, transactionID).Error
}

func (repository *TransactionRepositoryImpl) FindAll(ctx context.Context) (transactions []domain.Transaction, err error) {
	return transactions, repository.server.DB.WithContext(ctx).Find(&transactions).Error
}
