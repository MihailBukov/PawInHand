// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * PawInHand API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

type LoginUserRequest struct {
	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`
}

// AssertLoginUserRequestRequired checks if the required fields are not zero-ed
func AssertLoginUserRequestRequired(obj LoginUserRequest) error {
	return nil
}

// AssertLoginUserRequestConstraints checks if the values respects the defined constraints
func AssertLoginUserRequestConstraints(obj LoginUserRequest) error {
	return nil
}
