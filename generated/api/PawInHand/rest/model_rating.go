// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * PawInHand API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

type Rating struct {
	Score int32 `json:"score"`

	Comment string `json:"comment"`
}

// AssertRatingRequired checks if the required fields are not zero-ed
func AssertRatingRequired(obj Rating) error {
	elements := map[string]interface{}{
		"score":   obj.Score,
		"comment": obj.Comment,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRatingConstraints checks if the values respects the defined constraints
func AssertRatingConstraints(obj Rating) error {
	return nil
}
