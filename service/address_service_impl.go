package service

import (
	"context"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"

	"gorm.io/gorm"
)

type AddressServiceImpl struct {
	AddressRepository repository.AddressRepository
}

func NewAddressService(addressRepository repository.AddressRepository) AddressService {
	return &AddressServiceImpl{AddressRepository: addressRepository}
}

func (service *AddressServiceImpl) Create(ctx context.Context, request web.AddressRequest) (response web.AddressResponse, err error) {

	addressRequest := domain.Address{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		City:        request.City,
		ZipCode:     request.ZipCode,
		Country:     request.Country,
		UserID:      request.UserID,
	}

	addressResponse, err := service.AddressRepository.Create(ctx, addressRequest)
	if err != nil {
		return response, err
	}

	return web.AddressResponse{
		ID:          addressResponse.ID,
		FirstName:   addressResponse.FirstName,
		LastName:    addressResponse.LastName,
		PhoneNumber: addressResponse.PhoneNumber,
		Address:     addressResponse.Address,
		City:        addressResponse.City,
		ZipCode:     addressResponse.ZipCode,
		Country:     addressResponse.Country,
		UserID:      addressResponse.UserID,
	}, nil
}

func (service *AddressServiceImpl) Delete(ctx context.Context, addressID uint) error {
	var address domain.Address
	address.ID = addressID

	return service.AddressRepository.Delete(ctx, address)
}

func (service *AddressServiceImpl) FindByUser(ctx context.Context, userID uint) (responses []web.AddressResponse, err error) {
	addressResponse, err := service.AddressRepository.FindByUser(ctx, userID)

	if err != nil || len(addressResponse) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	for _, address := range addressResponse {
		responses = append(responses, web.AddressResponse{
			ID:          address.ID,
			FirstName:   address.FirstName,
			LastName:    address.LastName,
			PhoneNumber: address.PhoneNumber,
			Address:     address.Address,
			City:        address.City,
			ZipCode:     address.ZipCode,
			Country:     address.Country,
			UserID:      address.UserID,
		})
	}
	return responses, nil
}
