package service

import (
	"errors"
	"fmt"
	"public_service/adapter"
	"public_service/model"
)

type PublicService interface {
	Listings(page, size, userId int) ([]model.Listing, error)
	CreateListing(listing model.ListingRequest) (model.Listing, error)
	CreateUser(user model.UserRequest) (model.User, error)
}

type publicService struct {
	userAdapter    adapter.UserAdapter
	listingAdapter adapter.ListingAdapter
}

func NewPublicService(userAdapter adapter.UserAdapter, listingAdapter adapter.ListingAdapter) PublicService {
	return &publicService{
		userAdapter:    userAdapter,
		listingAdapter: listingAdapter,
	}
}

func (s *publicService) Listings(page, size, userId int) ([]model.Listing, error) {
	listingList, err := s.listingAdapter.GetListing(page, size, userId)
	if err != nil {
		return []model.Listing{}, errors.New("internal error")
	}

	res := make([]model.Listing, 0)
	for _, val := range listingList {
		curUser, err := s.userAdapter.GetUser(val.UserID)
		if err != nil {
			// do something
			continue
		}
		res = append(res, toListing(val, curUser))
	}

	return res, nil
}

func (s *publicService) CreateListing(listing model.ListingRequest) (model.Listing, error) {
	res, err := s.listingAdapter.CreateListing(listing)
	if err != nil {
		fmt.Println(err.Error())
		return model.Listing{}, errors.New("failed creating listing")
	}

	user, err := s.userAdapter.GetUser(res.UserID)
	if err != nil {
		return model.Listing{}, errors.New("user not found")
	}

	return toListing(res, user), nil
}

func (s *publicService) CreateUser(user model.UserRequest) (model.User, error) {
	return s.userAdapter.CreateUser(user)
}

func toListing(listingModel model.ListingModel, user model.User) model.Listing {
	return model.Listing{
		ID:          listingModel.ID,
		ListingType: listingModel.ListingType,
		Price:       listingModel.Price,
		CreatedAt:   listingModel.CreatedAt,
		UpdatedAt:   listingModel.UpdatedAt,
		User:        user,
	}
}
