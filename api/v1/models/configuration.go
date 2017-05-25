package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Configuration General purpose structure to hold configuration of the daemon and
// endpoints. Split into a mutable and immutable section.
//
// swagger:model Configuration
type Configuration struct {

	// Immutable configuration (read-only)
	Immutable ConfigurationMap `json:"immutable,omitempty"`

	// Changeable configuration
	Mutable ConfigurationMap `json:"mutable,omitempty"`
}

// Validate validates this configuration
func (m *Configuration) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configuration) UnmarshalBinary(b []byte) error {
	var res Configuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
