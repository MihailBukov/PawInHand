package resources

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/generated/dao/model"
	"log"
	"time"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/resources/mock_converter.go . ConverterI
type ConverterI interface {
	ConvertApiUserToModelUser(userApi *api.User) *model.User
	ConvertModelUserToApiUser(userModel *model.User) *api.User

	ConvertApiShelterToModelShelter(shelterApi *api.Shelter) *model.Shelter
	ConvertModelShelterToApiShelter(shelterModel *model.Shelter) *api.Shelter

	ConvertApiAdToModelAd(adApi *api.Ad) *model.Ad
	ConvertModelAdToApiAd(adModel *model.Ad) *api.Ad

	ConvertApiPostToModelPost(postApi *api.Post) *model.Post
	ConvertModelPostToApiPost(postModel *model.Post) *api.Post

	ConvertApiEventToModelEvent(eventApi *api.Event) *model.Event
	ConvertModelEventToApiEvent(eventModel *model.Event) *api.Event

	ConvertApiRatingToModelRating(ratingApi *api.Rating) *model.Rating
	ConvertModelRatingToApiRating(ratingModel *model.Rating) *api.Rating

	ConvertApiVolunteerWorkToModelVolunteerWork(workApi *api.VolunteerWork) *model.Volunteerwork
	ConvertModelVolunteerWorkToApiVolunteerWork(workModel *model.Volunteerwork) *api.VolunteerWork
}

type converter struct{}

func NewConverter() ConverterI {
	return &converter{}
}

// User
func (self *converter) ConvertApiUserToModelUser(userApi *api.User) *model.User {
	return &model.User{
		ID:        userApi.Id,
		FirstName: userApi.FirstName,
		LastName:  userApi.LastName,
		Email:     userApi.Email,
		Username:  userApi.Username,
	}
}

func (self *converter) ConvertModelUserToApiUser(userModel *model.User) *api.User {
	return &api.User{
		Id:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Username:  userModel.Username,
	}
}

// Shelter
func (self *converter) ConvertApiShelterToModelShelter(shelterApi *api.Shelter) *model.Shelter {
	return &model.Shelter{
		ID:     shelterApi.Id,
		Name:   shelterApi.Name,
		Phone:  shelterApi.Phone,
		Email:  shelterApi.Email,
		Street: shelterApi.Address.Street,
		City:   shelterApi.Address.City,
		State:  shelterApi.Address.State,
		Zip:    shelterApi.Address.Zip,
	}
}

func (self *converter) ConvertModelShelterToApiShelter(shelterModel *model.Shelter) *api.Shelter {
	return &api.Shelter{
		Id:    shelterModel.ID,
		Name:  shelterModel.Name,
		Phone: shelterModel.Phone,
		Email: shelterModel.Email,
		Address: api.ShelterAddress{
			Street: shelterModel.Street,
			City:   shelterModel.City,
			State:  shelterModel.State,
			Zip:    shelterModel.Zip,
		},
	}
}

// Ad
func (self *converter) ConvertApiAdToModelAd(adApi *api.Ad) *model.Ad {
	return &model.Ad{
		ID:          adApi.Id,
		Title:       adApi.Title,
		Description: adApi.Description,
		Image:       adApi.Image,
		DateCreated: time.Now(),
	}
}

func (self *converter) ConvertModelAdToApiAd(adModel *model.Ad) *api.Ad {
	return &api.Ad{
		Id:          adModel.ID,
		Title:       adModel.Title,
		Description: adModel.Description,
		Image:       adModel.Image,
		DateCreated: adModel.DateCreated.Format(time.RFC3339),
	}
}

// Post
func (self *converter) ConvertApiPostToModelPost(postApi *api.Post) *model.Post {
	return &model.Post{
		ID:          postApi.Id,
		Title:       postApi.Title,
		Content:     postApi.Content,
		AuthorID:    postApi.AuthorId,
		Image:       postApi.Image,
		DateCreated: time.Now(),
	}
}

func (self *converter) ConvertModelPostToApiPost(postModel *model.Post) *api.Post {
	return &api.Post{
		Id:          postModel.ID,
		Title:       postModel.Title,
		Content:     postModel.Content,
		AuthorId:    postModel.AuthorID,
		Image:       postModel.Image,
		DateCreated: postModel.DateCreated.Format(time.RFC3339),
	}
}

// Event
func (self *converter) ConvertApiEventToModelEvent(eventApi *api.Event) *model.Event {
	return &model.Event{
		ID:    eventApi.Id,
		Title: eventApi.Title,
		Date:  parseDate(eventApi.Date),
	}
}

func (self *converter) ConvertModelEventToApiEvent(eventModel *model.Event) *api.Event {
	return &api.Event{
		Id:    eventModel.ID,
		Title: eventModel.Title,
		Date:  eventModel.Date.Format(time.RFC3339),
	}
}

// Rating
func (self *converter) ConvertApiRatingToModelRating(ratingApi *api.Rating) *model.Rating {
	return &model.Rating{
		Score:   ratingApi.Score,
		Comment: ratingApi.Comment,
	}
}

func (self *converter) ConvertModelRatingToApiRating(ratingModel *model.Rating) *api.Rating {
	return &api.Rating{
		Score:   ratingModel.Score,
		Comment: ratingModel.Comment,
	}
}

// VolunteerWork
func (self *converter) ConvertApiVolunteerWorkToModelVolunteerWork(workApi *api.VolunteerWork) *model.Volunteerwork {
	return &model.Volunteerwork{
		ID:          workApi.Id,
		Title:       workApi.Title,
		Description: workApi.Description,
		Date:        parseDate(workApi.Date),
		Venue:       workApi.Location.Venue,
		City:        workApi.Location.City,
		State:       workApi.Location.State,
	}
}

func (self *converter) ConvertModelVolunteerWorkToApiVolunteerWork(workModel *model.Volunteerwork) *api.VolunteerWork {
	return &api.VolunteerWork{
		Id:          workModel.ID,
		Title:       workModel.Title,
		Description: workModel.Description,
		Date:        workModel.Date.Format(time.RFC3339),
		Location: api.VolunteerWorkLocation{
			Venue: workModel.Venue,
			City:  workModel.City,
			State: workModel.State,
		},
	}
}

// Helper
func parseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Println("Error parsing date:", err)
	}
	return date
}
