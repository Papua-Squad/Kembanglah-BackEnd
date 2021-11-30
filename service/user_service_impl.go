package service

import (
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SellerServiceImpl struct {
	SellerRepository repository.UserRepository
}

func NewSellerService(sellerRepository repository.UserRepository) UserService {
	return &SellerServiceImpl{
		SellerRepository: sellerRepository,
	}
}

func (service *SellerServiceImpl) Register(ctx context.Context, request web.SellerCreateRequest) web.Seller {

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	userRequest := domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hash),
	}
	user, err := service.SellerRepository.Create(ctx, userRequest)
	helper.PanicIfError(err)

	return web.Seller{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func (service *SellerServiceImpl) Update(ctx context.Context, request web.SellerUpdateRequest) web.Seller {
	var seller domain.User
	seller.ID = request.ID

	_, err := service.SellerRepository.FindByID(ctx, seller.ID)
	helper.PanicIfError(err)

	user, err := service.SellerRepository.Update(ctx, domain.User{
		Model: gorm.Model{
			ID: request.ID,
		},
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
	})
	helper.PanicIfError(err)

	return web.Seller{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func (service *SellerServiceImpl) Delete(ctx context.Context, sellerId uint) {
	var seller domain.User
	seller.ID = sellerId

	err := service.SellerRepository.Delete(ctx, seller)
	helper.PanicIfError(err)
}

func (service *SellerServiceImpl) FindByID(ctx context.Context, sellerId uint) web.Seller {
	user, err := service.SellerRepository.FindByID(ctx, sellerId)
	helper.PanicIfError(err)

	return web.Seller{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func (service *SellerServiceImpl) FindByUsername(ctx context.Context, username string) web.Seller {
	user, err := service.SellerRepository.FindByUsername(ctx, username)
	helper.PanicIfError(err)

	return web.Seller{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func (service *SellerServiceImpl) FindAll(ctx context.Context) []web.Seller {
	var newSeller []web.Seller

	seller, err := service.SellerRepository.FindAll(ctx)
	helper.PanicIfError(err)

	for _, data := range seller {
		newSeller = append(newSeller, web.Seller{
			ID:        data.ID,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Username:  data.Username,
			Email:     data.Email,
			Password:  data.Password,
		})
	}
	return newSeller

}

// func (service *UserServiceImpl) Login(ctx context.Context, userId uint) web.User {
// 	user := domain.User{}
// 	user.ID = userId

// 	user, err := service.UserRepository.FindByID(ctx, user)
// 	helper.PanicIfError(err)

// 	return web.User{
// 		ID: user.ID,
// 	}
// }
