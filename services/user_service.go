package services

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/repositories"
	"PawInHand/resources"
	"context"
	"net/http"
)

type UserAPIService struct {
	DB      repositories.UserRepo
	Convert resources.ConverterI
}

func NewUserAPIService(db repositories.UserRepo, convert resources.ConverterI) api.UserAPIServicer {
	return &UserAPIService{
		DB:      db,
		Convert: convert,
	}
}

// GetAllUsers -
func (s *UserAPIService) GetAllUsers(ctx context.Context) (api.ImplResponse, error) {
	users, err := s.DB.GetAll()
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	var userApis []*api.User
	for _, user := range users {
		userApi := s.Convert.ConvertModelUserToApiUser(user)
		userApis = append(userApis, userApi)
	}

	return api.Response(http.StatusOK, userApis), nil
}

// GetUserById -
func (s *UserAPIService) GetUserById(ctx context.Context, userId string) (api.ImplResponse, error) {
	if userId == "" {
		return api.Response(http.StatusBadRequest, "userId is required"), nil
	}

	user, err := s.DB.FindByID(userId)
	if err != nil {
		return api.Response(http.StatusNotFound, "User not found"), err
	}

	userApi := s.Convert.ConvertModelUserToApiUser(user)
	return api.Response(http.StatusOK, userApi), nil
}

// RegisterUser -
func (s *UserAPIService) RegisterUser(ctx context.Context, user api.User) (api.ImplResponse, error) {
	userModel := s.Convert.ConvertApiUserToModelUser(&user)
	newUser, err := s.DB.Register(userModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	userApi := s.Convert.ConvertModelUserToApiUser(newUser)
	return api.Response(http.StatusCreated, userApi), nil
}

// DeleteUser -
func (s *UserAPIService) DeleteUser(ctx context.Context, userId string) (api.ImplResponse, error) {
	if userId == "" {
		return api.Response(http.StatusBadRequest, "userId is required"), nil
	}

	_, err := s.DB.DeleteByID(userId)
	if err != nil {
		return api.Response(http.StatusNotFound, "User not found"), err
	}

	return api.Response(http.StatusNoContent, nil), nil
}
