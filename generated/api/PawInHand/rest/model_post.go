// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * PawInHand API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

type Post struct {
	Id string `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Content string `json:"content,omitempty"`

	AuthorId string `json:"authorId,omitempty"`

	DateCreated string `json:"dateCreated,omitempty"`

	Image string `json:"image,omitempty"`
}

// AssertPostRequired checks if the required fields are not zero-ed
func AssertPostRequired(obj Post) error {
	return nil
}

// AssertPostConstraints checks if the values respects the defined constraints
func AssertPostConstraints(obj Post) error {
	return nil
}
