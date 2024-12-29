// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * PawInHand API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

type VolunteerWorkLocation struct {
	Venue string `json:"venue,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`
}

// AssertVolunteerWorkLocationRequired checks if the required fields are not zero-ed
func AssertVolunteerWorkLocationRequired(obj VolunteerWorkLocation) error {
	return nil
}

// AssertVolunteerWorkLocationConstraints checks if the values respects the defined constraints
func AssertVolunteerWorkLocationConstraints(obj VolunteerWorkLocation) error {
	return nil
}
