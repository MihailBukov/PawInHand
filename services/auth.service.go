package services

//
//import (
//	api "PawInHand/generated/api/PawInHand/rest"
//	"PawInHand/repositories"
//	"context"
//	"net/http"
//)
//
//type authAPIService struct {
//	DB repositories.AuthRepo
//}
//
//func NewAuthAPIService(db repositories.AuthRepo) api.AuthAPIServicer {
//	return &authAPIService{
//		DB: db,
//	}
//}
//
//func (s *authAPIService) LoginUser(ctx context.Context, loginRequest api.LoginUserRequest) (api.ImplResponse, error) {
//	if loginRequest.Email == "" || loginRequest.Password == "" {
//		return api.Response(http.StatusBadRequest, "Email and password are required"), nil
//	}
//
//	user, err := s.DB.Login(loginRequest.Email, loginRequest.Password)
//	if err != nil {
//		return api.Response(http.StatusUnauthorized, "Invalid credentials"), err
//	}
//
//	// Generate token logic would be here
//	token := "generated-jwt-token" // Placeholder for actual token generation
//
//	return api.Response(http.StatusOK, api.LoginUser200Response{Token: token}), nil
//}
//
//func (s *authAPIService) LogoutUser(ctx context.Context, userID string) (api.ImplResponse, error) {
//	if userID == "" {
//		return api.Response(http.StatusBadRequest, "User ID is required"), nil
//	}
//
//	err := s.DB.Logout(userID)
//	if err != nil {
//		return api.Response(http.StatusInternalServerError, "Failed to logout"), err
//	}
//
//	return api.Response(http.StatusOK, "Logout successful"), nil
//}
