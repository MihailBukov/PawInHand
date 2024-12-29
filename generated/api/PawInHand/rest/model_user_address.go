// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * PawInHand API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

type UserAddress struct {
	Street string `json:"street,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Zip string `json:"zip,omitempty"`
}

// AssertUserAddressRequired checks if the required fields are not zero-ed
func AssertUserAddressRequired(obj UserAddress) error {
	return nil
}

// AssertUserAddressConstraints checks if the values respects the defined constraints
func AssertUserAddressConstraints(obj UserAddress) error {
	return nil
}