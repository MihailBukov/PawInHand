package services

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/generated/dao/model"
	"PawInHand/repositories"
	"context"
	"net/http"
)

type RatingAPIService struct {
	DB repositories.RatingRepo
}

func NewRatingAPIService(db repositories.RatingRepo) api.RatingAPIServicer {
	return &RatingAPIService{
		DB: db,
	}
}

// SubmitRating -
func (s *RatingAPIService) RateUser(ctx context.Context, userId string, rating api.Rating) (api.ImplResponse, error) {
	ratingModel := &model.Rating{
		UserID:  userId,
		Score:   rating.Score,
		Comment: rating.Comment,
	}
	newRating, err := s.DB.Create(ratingModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusCreated, newRating), nil
}
