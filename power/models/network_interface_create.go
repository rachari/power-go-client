// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkInterfaceCreate network interface create
//
// swagger:model NetworkInterfaceCreate
type NetworkInterfaceCreate struct {

	// The requested IP address of this Network Interface
	IPAddress string `json:"ipAddress,omitempty"`

	// Name of the Network Interface
	Name string `json:"name,omitempty"`

	// The user tags associated with this resource.
	UserTags []string `json:"userTags"`
}

// Validate validates this network interface create
func (m *NetworkInterfaceCreate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network interface create based on context it is used
func (m *NetworkInterfaceCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfaceCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfaceCreate) UnmarshalBinary(b []byte) error {
	var res NetworkInterfaceCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
