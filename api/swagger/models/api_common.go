// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// APICommon api common
// swagger:model apiCommon
type APICommon struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this api common
func (m *APICommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APICommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICommon) UnmarshalBinary(b []byte) error {
	var res APICommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}