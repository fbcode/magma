// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// S6a s6a configuration
// swagger:model s6a
type S6a struct {

	// server
	Server *DiameterClientConfigs `json:"server,omitempty"`
}

// Validate validates this s6a
func (m *S6a) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S6a) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S6a) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S6a) UnmarshalBinary(b []byte) error {
	var res S6a
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
