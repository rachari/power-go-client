// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IKEPolicyCreate IKE Policy object used for creation
//
// swagger:model IKEPolicyCreate
type IKEPolicyCreate struct {

	// authentication
	Authentication IKEPolicyAuthentication `json:"authentication,omitempty"`

	// DH group of the IKE Policy
	// Example: 2
	// Required: true
	// Enum: [1,2,5,14,19,20,24]
	DhGroup *int64 `json:"dhGroup"`

	// encryption of the IKE Policy
	// Example: aes-256-cbc
	// Required: true
	// Enum: ["aes-256-cbc","aes-192-cbc","aes-128-cbc","aes-256-gcm","aes-128-gcm","3des-cbc"]
	Encryption *string `json:"encryption"`

	// key lifetime
	// Required: true
	KeyLifetime *KeyLifetime `json:"keyLifetime"`

	// name of the IKE Policy
	// Example: ikePolicy1
	// Required: true
	// Max Length: 47
	// Min Length: 1
	Name *string `json:"name"`

	// Preshared key used in this IKE Policy (length of preshared key must be even)
	// Required: true
	PresharedKey *string `json:"presharedKey"`

	// version of the IKE Policy
	// Example: 2
	// Required: true
	// Enum: [1,2]
	Version *int64 `json:"version"`
}

// Validate validates this i k e policy create
func (m *IKEPolicyCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresharedKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IKEPolicyCreate) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if err := m.Authentication.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authentication")
		}
		return err
	}

	return nil
}

var iKEPolicyCreateTypeDhGroupPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,5,14,19,20,24]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyCreateTypeDhGroupPropEnum = append(iKEPolicyCreateTypeDhGroupPropEnum, v)
	}
}

// prop value enum
func (m *IKEPolicyCreate) validateDhGroupEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, iKEPolicyCreateTypeDhGroupPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicyCreate) validateDhGroup(formats strfmt.Registry) error {

	if err := validate.Required("dhGroup", "body", m.DhGroup); err != nil {
		return err
	}

	// value enum
	if err := m.validateDhGroupEnum("dhGroup", "body", *m.DhGroup); err != nil {
		return err
	}

	return nil
}

var iKEPolicyCreateTypeEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aes-256-cbc","aes-192-cbc","aes-128-cbc","aes-256-gcm","aes-128-gcm","3des-cbc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyCreateTypeEncryptionPropEnum = append(iKEPolicyCreateTypeEncryptionPropEnum, v)
	}
}

const (

	// IKEPolicyCreateEncryptionAesDash256DashCbc captures enum value "aes-256-cbc"
	IKEPolicyCreateEncryptionAesDash256DashCbc string = "aes-256-cbc"

	// IKEPolicyCreateEncryptionAesDash192DashCbc captures enum value "aes-192-cbc"
	IKEPolicyCreateEncryptionAesDash192DashCbc string = "aes-192-cbc"

	// IKEPolicyCreateEncryptionAesDash128DashCbc captures enum value "aes-128-cbc"
	IKEPolicyCreateEncryptionAesDash128DashCbc string = "aes-128-cbc"

	// IKEPolicyCreateEncryptionAesDash256DashGcm captures enum value "aes-256-gcm"
	IKEPolicyCreateEncryptionAesDash256DashGcm string = "aes-256-gcm"

	// IKEPolicyCreateEncryptionAesDash128DashGcm captures enum value "aes-128-gcm"
	IKEPolicyCreateEncryptionAesDash128DashGcm string = "aes-128-gcm"

	// IKEPolicyCreateEncryptionNr3desDashCbc captures enum value "3des-cbc"
	IKEPolicyCreateEncryptionNr3desDashCbc string = "3des-cbc"
)

// prop value enum
func (m *IKEPolicyCreate) validateEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iKEPolicyCreateTypeEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicyCreate) validateEncryption(formats strfmt.Registry) error {

	if err := validate.Required("encryption", "body", m.Encryption); err != nil {
		return err
	}

	// value enum
	if err := m.validateEncryptionEnum("encryption", "body", *m.Encryption); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicyCreate) validateKeyLifetime(formats strfmt.Registry) error {

	if err := validate.Required("keyLifetime", "body", m.KeyLifetime); err != nil {
		return err
	}

	if err := validate.Required("keyLifetime", "body", m.KeyLifetime); err != nil {
		return err
	}

	if m.KeyLifetime != nil {
		if err := m.KeyLifetime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyLifetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyLifetime")
			}
			return err
		}
	}

	return nil
}

func (m *IKEPolicyCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 47); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicyCreate) validatePresharedKey(formats strfmt.Registry) error {

	if err := validate.Required("presharedKey", "body", m.PresharedKey); err != nil {
		return err
	}

	return nil
}

var iKEPolicyCreateTypeVersionPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyCreateTypeVersionPropEnum = append(iKEPolicyCreateTypeVersionPropEnum, v)
	}
}

// prop value enum
func (m *IKEPolicyCreate) validateVersionEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, iKEPolicyCreateTypeVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicyCreate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	// value enum
	if err := m.validateVersionEnum("version", "body", *m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this i k e policy create based on the context it is used
func (m *IKEPolicyCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyLifetime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IKEPolicyCreate) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authentication")
		}
		return err
	}

	return nil
}

func (m *IKEPolicyCreate) contextValidateKeyLifetime(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyLifetime != nil {

		if err := m.KeyLifetime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyLifetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyLifetime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IKEPolicyCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IKEPolicyCreate) UnmarshalBinary(b []byte) error {
	var res IKEPolicyCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
