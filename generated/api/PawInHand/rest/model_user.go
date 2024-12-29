// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * PawInHand API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

type User struct {
	Id string `json:"id,omitempty"`

	FirstName string `json:"firstName"`

	MiddleName string `json:"middleName,omitempty"`

	LastName string `json:"lastName"`

	Username string `json:"username"`

	Email string `json:"email"`

	Phone string `json:"phone,omitempty"`

	Address UserAddress `json:"address,omitempty"`

	Role string `json:"role,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"firstName": obj.FirstName,
		"lastName":  obj.LastName,
		"username":  obj.Username,
		"email":     obj.Email,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUserAddressRequired(obj.Address); err != nil {
		return err
	}
	return nil
}

// AssertUserConstraints checks if the values respects the defined constraints
func AssertUserConstraints(obj User) error {
	if err := AssertUserAddressConstraints(obj.Address); err != nil {
		return err
	}
	return nil
}
