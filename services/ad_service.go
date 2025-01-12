package services

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/generated/dao/model"
	"PawInHand/repositories"
	"PawInHand/resources"
	"context"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type AdAPIService struct {
	DB      repositories.AdRepo
	Convert resources.ConverterI
}

func NewAdAPIService(db repositories.AdRepo, convert resources.ConverterI) api.AdsAPIServicer {
	return &AdAPIService{
		DB:      db,
		Convert: convert,
	}
}

// GetAllAds -
func (s *AdAPIService) GetAllAds(ctx context.Context) (api.ImplResponse, error) {
	ads, err := s.DB.GetAll()
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	var adApis []*api.Ad
	for _, ad := range ads {
		adApi := s.Convert.ConvertModelAdToApiAd(ad)
		adApis = append(adApis, adApi)
	}

	return api.Response(http.StatusOK, adApis), nil
}

// GetAdById -
func (s *AdAPIService) GetAdById(ctx context.Context, adId string) (api.ImplResponse, error) {
	if adId == "" {
		return api.Response(http.StatusBadRequest, "adId is required"), nil
	}

	ad, err := s.DB.FindByID(adId)
	if err != nil {
		return api.Response(http.StatusNotFound, "Ad not found"), err
	}

	adApi := s.Convert.ConvertModelAdToApiAd(ad)
	return api.Response(http.StatusOK, adApi), nil
}

// CreateAd -
func (s *AdAPIService) CreateAd(ctx context.Context, ad api.Ad) (api.ImplResponse, error) {
	adModel := s.Convert.ConvertApiAdToModelAd(&ad)
	newAd, err := s.DB.Create(adModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	adApi := s.Convert.ConvertModelAdToApiAd(newAd)
	return api.Response(http.StatusCreated, adApi), nil
}

// DeleteAd -
func (s *AdAPIService) DeleteAdById(ctx context.Context, adId string) (api.ImplResponse, error) {
	if adId == "" {
		return api.Response(http.StatusBadRequest, "adId is required"), nil
	}

	_, err := s.DB.DeleteByID(adId)
	if err != nil {
		return api.Response(http.StatusNotFound, "Ad not found"), err
	}

	return api.Response(http.StatusNoContent, nil), nil
}
func (s *AdAPIService) UpdateAdById(ctx context.Context, adID string, ad api.Ad) (api.ImplResponse, error) {
	// Validate input: Ensure required fields are present
	if ad.Title == "" || ad.Description == "" {
		return api.Response(http.StatusBadRequest, "Title and description are required"), nil
	}

	// Parse and validate date fields if applicable
	var adDate time.Time
	var err error
	if ad.DateCreated != "" {
		adDate, err = time.Parse("2006-01-02", ad.DateCreated)
		if err != nil {
			return api.Response(http.StatusBadRequest, "Invalid date format"), nil
		}
	}

	// Map API model to database/DAO model
	updatedAd := &model.Ad{
		ID:          adID,
		Title:       ad.Title,
		Description: ad.Description,
		DateCreated: adDate,
		Image:       ad.Image,
	}

	// Update the ad in the database
	_, err = s.DB.UpdateById(updatedAd)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return api.Response(http.StatusNotFound, "Ad not found"), nil
		}
		return api.Response(http.StatusInternalServerError, "Failed to update ad"), err
	}

	// Return success response
	return api.Response(http.StatusOK, "Ad updated successfully"), nil
}
