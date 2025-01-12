package services

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/generated/dao/model"
	"PawInHand/repositories"
	"PawInHand/resources"
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type ShelterAPIService struct {
	DB      repositories.ShelterRepo
	Convert resources.ConverterI
}

func NewShelterAPIService(db repositories.ShelterRepo, convert resources.ConverterI) api.ShelterAPIServicer {
	return &ShelterAPIService{
		DB:      db,
		Convert: convert,
	}
}

// GetAllShelters -
func (s *ShelterAPIService) GetAllShelters(ctx context.Context) (api.ImplResponse, error) {
	shelters, err := s.DB.GetAll()
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	var shelterApis []*api.Shelter
	for _, shelter := range shelters {
		shelterApi := s.Convert.ConvertModelShelterToApiShelter(shelter)
		shelterApis = append(shelterApis, shelterApi)
	}

	return api.Response(http.StatusOK, shelterApis), nil
}

// GetShelterById -
func (s *ShelterAPIService) GetShelterById(ctx context.Context, shelterId string) (api.ImplResponse, error) {
	if shelterId == "" {
		return api.Response(http.StatusBadRequest, "shelterId is required"), nil
	}

	shelter, err := s.DB.FindByID(shelterId)
	if err != nil {
		return api.Response(http.StatusNotFound, "Shelter not found"), err
	}

	shelterApi := s.Convert.ConvertModelShelterToApiShelter(shelter)
	return api.Response(http.StatusOK, shelterApi), nil
}

// RegisterShelter -
func (s *ShelterAPIService) RegisterShelter(ctx context.Context, shelter api.Shelter) (api.ImplResponse, error) {
	shelterModel := s.Convert.ConvertApiShelterToModelShelter(&shelter)
	newShelter, err := s.DB.Register(shelterModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	shelterApi := s.Convert.ConvertModelShelterToApiShelter(newShelter)
	return api.Response(http.StatusCreated, shelterApi), nil
}

// DeleteShelter -
func (s *ShelterAPIService) DeleteShelterById(ctx context.Context, shelterId string) (api.ImplResponse, error) {
	if shelterId == "" {
		return api.Response(http.StatusBadRequest, "shelterId is required"), nil
	}

	_, err := s.DB.DeleteByID(shelterId)
	if err != nil {
		return api.Response(http.StatusNotFound, "Shelter not found"), err
	}

	return api.Response(http.StatusNoContent, nil), nil
}

func (s *ShelterAPIService) UpdateShelterById(ctx context.Context, shelterId string, shelter api.Shelter) (api.ImplResponse, error) {
	// Validate input
	if shelterId == "" {
		return api.Response(http.StatusBadRequest, "Shelter ID is required"), nil
	}
	if shelter.Name == "" || shelter.Address.Zip == "" {
		return api.Response(http.StatusBadRequest, "Shelter name and address are required"), nil
	}

	// Convert API model to database model
	shelterModel := &model.Shelter{
		ID:        shelterId,
		Name:      shelter.Name,
		Street:    shelter.Address.Street,
		Zip:       shelter.Address.Zip,
		State:     shelter.Address.State,
		City:      shelter.Address.City,
		UpdatedAt: time.Now(),
		Phone:     shelter.Phone,
		Email:     shelter.Email,
	}

	// Perform the update operation using the repository
	updatedShelter, err := s.DB.UpdateByID(shelterId, shelterModel)
	if err != nil {
		// Handle database errors
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api.Response(http.StatusNotFound, "Shelter not found"), nil
		}
		return api.Response(http.StatusInternalServerError, "Failed to update shelter"), err
	}

	// Return success response
	return api.Response(http.StatusOK, updatedShelter), nil
}
