package service

import (
	"context"
	"gorm.io/gorm"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"
)

type TransactionServiceImpl struct {
	ProductRepository     repository.ProductRepository
	TransactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductRepository) TransactionService {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository, ProductRepository: productRepository}
}

func (service *TransactionServiceImpl) Save(ctx context.Context, request web.CreateTransactionRequest) (response web.TransactionResponse, err error) {
	var productsId []int64
	var products []web.ProductResponse
	var subTotal int
	var qty int
	for _, id := range request.ProductsID {
		productsId = append(productsId, int64(id))
		product, _ := service.ProductRepository.FindByID(ctx, uint(id))
		products = append(products, web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Weight:      product.Weight,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			SellerID:    product.SellerID,
			ImageUrl:    product.ImageUrl,
		})

		subTotal += product.Price
		qty += 1
	}
	transactionRequest := domain.Transaction{
		ProductsID:      productsId,
		SellerID:        request.SellerID,
		CustomerID:      request.CustomerID,
		ShippingAddress: request.ShippingAddress,
		PaymentMethode:  request.PaymentMethode,
		Status:          request.Status,
		SubTotal:        subTotal,
		Qty:             qty,
	}

	transactionResponse, err := service.TransactionRepository.Save(ctx, transactionRequest)
	if err != nil {
		return response, err
	}

	return web.TransactionResponse{
		ID:              transactionResponse.ID,
		Products:        products,
		SellerID:        transactionResponse.SellerID,
		CustomerID:      transactionResponse.CustomerID,
		ShippingAddress: transactionResponse.ShippingAddress,
		PaymentMethode:  transactionResponse.PaymentMethode,
		SubTotal:        transactionResponse.SubTotal,
		Qty:             transactionResponse.Qty,
		Status:          transactionResponse.Status,
	}, nil

}

func (service *TransactionServiceImpl) Update(ctx context.Context, request web.UpdateStatusTransactionRequest) (response web.TransactionResponse, err error) {
	var transaction domain.Transaction
	transaction.ID = request.ID

	_, err = service.TransactionRepository.FindByID(ctx, transaction.ID)
	if err != nil {
		return response, err
	}

	transactionResponse, err := service.TransactionRepository.Update(ctx, domain.Transaction{
		Model:  gorm.Model{ID: transaction.ID},
		Status: request.Status,
	})
	if err != nil {
		return response, err
	}

	// Init products
	var products []web.ProductResponse

	for _, id := range transactionResponse.ProductsID {
		product, _ := service.ProductRepository.FindByID(ctx, uint(id))
		products = append(products, web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Weight:      product.Weight,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			SellerID:    product.SellerID,
			ImageUrl:    product.ImageUrl,
		})

	}

	return web.TransactionResponse{
		ID:              transactionResponse.ID,
		Products:        products,
		SellerID:        transactionResponse.SellerID,
		CustomerID:      transactionResponse.CustomerID,
		ShippingAddress: transactionResponse.ShippingAddress,
		PaymentMethode:  transactionResponse.PaymentMethode,
		SubTotal:        transactionResponse.SubTotal,
		Qty:             transactionResponse.Qty,
		Status:          transactionResponse.Status,
	}, nil
}

func (service *TransactionServiceImpl) Delete(ctx context.Context, TransactionID uint) error {
	var transaction domain.Transaction
	transaction.ID = TransactionID

	return service.TransactionRepository.Delete(ctx, transaction)
}

func (service *TransactionServiceImpl) FindBySeller(ctx context.Context, sellerID uint) (responses []web.TransactionResponse, err error) {
	transactionResponse, err := service.TransactionRepository.FindBySeller(ctx, sellerID)
	if err != nil {
		return responses, gorm.ErrRecordNotFound
	}

	for _, transaction := range transactionResponse {
		var products []web.ProductResponse
		for _, id := range transaction.ProductsID {

			product, _ := service.ProductRepository.FindByID(ctx, uint(id))
			products = append(products, web.ProductResponse{
				ID:          product.ID,
				Name:        product.Name,
				Price:       product.Price,
				Stock:       product.Stock,
				Weight:      product.Weight,
				Description: product.Description,
				CategoryID:  product.CategoryID,
				SellerID:    product.SellerID,
				ImageUrl:    product.ImageUrl,
			})
		}

		responses = append(responses, web.TransactionResponse{
			ID:              transaction.ID,
			Products:        products,
			SellerID:        transaction.SellerID,
			CustomerID:      transaction.CustomerID,
			ShippingAddress: transaction.ShippingAddress,
			PaymentMethode:  transaction.PaymentMethode,
			SubTotal:        transaction.SubTotal,
			Qty:             transaction.Qty,
			Status:          transaction.Status,
		})
		// reset
		products = nil
	}

	return responses, nil
}

func (service *TransactionServiceImpl) FindByCustomer(ctx context.Context, customerID uint) (responses []web.TransactionResponse, err error) {
	transactionResponse, err := service.TransactionRepository.FindByCustomer(ctx, customerID)
	if err != nil {
		return responses, gorm.ErrRecordNotFound
	}

	for _, transaction := range transactionResponse {
		var products []web.ProductResponse
		for _, id := range transaction.ProductsID {

			product, _ := service.ProductRepository.FindByID(ctx, uint(id))
			products = append(products, web.ProductResponse{
				ID:          product.ID,
				Name:        product.Name,
				Price:       product.Price,
				Stock:       product.Stock,
				Weight:      product.Weight,
				Description: product.Description,
				CategoryID:  product.CategoryID,
				SellerID:    product.SellerID,
				ImageUrl:    product.ImageUrl,
			})
		}

		responses = append(responses, web.TransactionResponse{
			ID:              transaction.ID,
			Products:        products,
			SellerID:        transaction.SellerID,
			CustomerID:      transaction.CustomerID,
			ShippingAddress: transaction.ShippingAddress,
			PaymentMethode:  transaction.PaymentMethode,
			SubTotal:        transaction.SubTotal,
			Qty:             transaction.Qty,
			Status:          transaction.Status,
		})
		// reset
		products = nil
	}

	return responses, nil
}

func (service *TransactionServiceImpl) FindByID(ctx context.Context, transactionID uint) (response web.TransactionResponse, err error) {
	var products []web.ProductResponse

	transactionResponse, err := service.TransactionRepository.FindByID(ctx, transactionID)
	if err != nil {
		return response, gorm.ErrRecordNotFound
	}

	for _, id := range transactionResponse.ProductsID {
		product, _ := service.ProductRepository.FindByID(ctx, uint(id))
		products = append(products, web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Weight:      product.Weight,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			SellerID:    product.SellerID,
			ImageUrl:    product.ImageUrl,
		})

	}

	return web.TransactionResponse{
		ID:              transactionResponse.ID,
		Products:        products,
		SellerID:        transactionResponse.SellerID,
		CustomerID:      transactionResponse.CustomerID,
		ShippingAddress: transactionResponse.ShippingAddress,
		PaymentMethode:  transactionResponse.PaymentMethode,
		SubTotal:        transactionResponse.SubTotal,
		Qty:             transactionResponse.Qty,
		Status:          transactionResponse.Status,
	}, nil

}

func (service *TransactionServiceImpl) FindAll(ctx context.Context) (responses []web.TransactionResponse, err error) {
	transactionResponse, err := service.TransactionRepository.FindAll(ctx)
	if err != nil {
		return responses, gorm.ErrRecordNotFound
	}

	for _, transaction := range transactionResponse {
		var products []web.ProductResponse
		for _, id := range transaction.ProductsID {

			product, _ := service.ProductRepository.FindByID(ctx, uint(id))
			products = append(products, web.ProductResponse{
				ID:          product.ID,
				Name:        product.Name,
				Price:       product.Price,
				Stock:       product.Stock,
				Weight:      product.Weight,
				Description: product.Description,
				CategoryID:  product.CategoryID,
				SellerID:    product.SellerID,
				ImageUrl:    product.ImageUrl,
			})
		}

		responses = append(responses, web.TransactionResponse{
			ID:              transaction.ID,
			Products:        products,
			SellerID:        transaction.SellerID,
			CustomerID:      transaction.CustomerID,
			ShippingAddress: transaction.ShippingAddress,
			PaymentMethode:  transaction.PaymentMethode,
			SubTotal:        transaction.SubTotal,
			Qty:             transaction.Qty,
			Status:          transaction.Status,
		})
		// reset
		products = nil
	}

	return responses, nil

}
