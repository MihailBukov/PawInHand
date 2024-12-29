// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * PawInHand API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// VolunteerAPIController binds http requests to an api service and writes the service results to the http response
type VolunteerAPIController struct {
	service      VolunteerAPIServicer
	errorHandler ErrorHandler
}

// VolunteerAPIOption for how the controller is set up.
type VolunteerAPIOption func(*VolunteerAPIController)

// WithVolunteerAPIErrorHandler inject ErrorHandler into controller
func WithVolunteerAPIErrorHandler(h ErrorHandler) VolunteerAPIOption {
	return func(c *VolunteerAPIController) {
		c.errorHandler = h
	}
}

// NewVolunteerAPIController creates a default api controller
func NewVolunteerAPIController(s VolunteerAPIServicer, opts ...VolunteerAPIOption) *VolunteerAPIController {
	controller := &VolunteerAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the VolunteerAPIController
func (c *VolunteerAPIController) Routes() Routes {
	return Routes{
		"CreateVolunteerWork": Route{
			strings.ToUpper("Post"),
			"/api/volunteer-work",
			c.CreateVolunteerWork,
		},
		"DeleteVolunteerWorkById": Route{
			strings.ToUpper("Delete"),
			"/api/volunteer-work/{workId}",
			c.DeleteVolunteerWorkById,
		},
		"GetAllVolunteerWork": Route{
			strings.ToUpper("Get"),
			"/api/volunteer-work",
			c.GetAllVolunteerWork,
		},
		"GetVolunteerWorkById": Route{
			strings.ToUpper("Get"),
			"/api/volunteer-work/{workId}",
			c.GetVolunteerWorkById,
		},
		"RegisterForVolunteerWork": Route{
			strings.ToUpper("Post"),
			"/api/volunteer-work/{workId}/register",
			c.RegisterForVolunteerWork,
		},
	}
}

// CreateVolunteerWork -
func (c *VolunteerAPIController) CreateVolunteerWork(w http.ResponseWriter, r *http.Request) {
	volunteerWorkParam := VolunteerWork{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&volunteerWorkParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertVolunteerWorkRequired(volunteerWorkParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertVolunteerWorkConstraints(volunteerWorkParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateVolunteerWork(r.Context(), volunteerWorkParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteVolunteerWorkById -
func (c *VolunteerAPIController) DeleteVolunteerWorkById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workIdParam := params["workId"]
	if workIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"workId"}, nil)
		return
	}
	result, err := c.service.DeleteVolunteerWorkById(r.Context(), workIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAllVolunteerWork -
func (c *VolunteerAPIController) GetAllVolunteerWork(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAllVolunteerWork(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetVolunteerWorkById -
func (c *VolunteerAPIController) GetVolunteerWorkById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workIdParam := params["workId"]
	if workIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"workId"}, nil)
		return
	}
	result, err := c.service.GetVolunteerWorkById(r.Context(), workIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// RegisterForVolunteerWork -
func (c *VolunteerAPIController) RegisterForVolunteerWork(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workIdParam := params["workId"]
	if workIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"workId"}, nil)
		return
	}
	volunteerRegistrationParam := VolunteerRegistration{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&volunteerRegistrationParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertVolunteerRegistrationRequired(volunteerRegistrationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertVolunteerRegistrationConstraints(volunteerRegistrationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RegisterForVolunteerWork(r.Context(), workIdParam, volunteerRegistrationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}