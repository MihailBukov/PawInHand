package services

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/generated/dao/model"
	"PawInHand/repositories"
	"context"
	"net/http"
	"time"
)

type VolunteerWorkAPIService struct {
	DB repositories.VolunteerworkRepo
}

func NewVolunteerWorkAPIService(db repositories.VolunteerworkRepo) api.VolunteerAPIServicer {
	return &VolunteerWorkAPIService{
		DB: db,
	}
}

// RegisterForWork -
func (s *VolunteerWorkAPIService) RegisterForVolunteerWork(ctx context.Context, workId string, work api.VolunteerWork) (api.ImplResponse, error) {
	date, err := time.Parse("2006-01-02", work.Date)

	workModel := &model.Volunteerwork{
		ID:          workId,
		Title:       work.Title,
		Description: work.Description,
		Date:        date,
	}
	newWork, err := s.DB.Register(workModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusCreated, newWork), nil
}
