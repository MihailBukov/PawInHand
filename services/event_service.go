package services

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/generated/dao/model"
	"PawInHand/repositories"
	"PawInHand/resources"
	"context"
	"net/http"
	"time"
)

type EventAPIService struct {
	DB        repositories.EventRepo
	Converter resources.ConverterI
}

func NewEventAPIService(db repositories.EventRepo) api.EventsAPIServicer {
	return &EventAPIService{
		DB: db,
	}
}

// GetEvents -
func (s *EventAPIService) GetEvents(ctx context.Context) (api.ImplResponse, error) {
	events, err := s.DB.GetAll()
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusOK, events), nil
}

// CreateEvent -
func (s *EventAPIService) CreateEvent(ctx context.Context, event api.Event) (api.ImplResponse, error) {
	date, err := time.Parse("2006-01-02", event.Date)

	eventModel := &model.Event{
		ID:    event.Id,
		Title: event.Title,
		Date:  date,
	}
	newEvent, err := s.DB.Create(eventModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusCreated, newEvent), nil
}
