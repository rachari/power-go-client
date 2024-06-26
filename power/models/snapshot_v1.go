// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotV1 snapshot v1
//
// swagger:model SnapshotV1
type SnapshotV1 struct {

	// The date and time when the snapshot was created.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The snapshot UUID.
	// Required: true
	ID *string `json:"id"`

	// The snapshot name.
	// Required: true
	Name *string `json:"name"`

	// The size of the snapshot, in gibibytes (GiB).
	// Required: true
	Size *float64 `json:"size"`

	// The status for the snapshot.
	// Required: true
	Status *string `json:"status"`

	// The date and time when the snapshot was last updated.
	// Format: date-time
	UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`

	// The volume UUID associated with the snapshot.
	// Required: true
	VolumeID *string `json:"volumeID"`
}

// Validate validates this snapshot v1
func (m *SnapshotV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotV1) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotV1) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotV1) validateUpdatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDate", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotV1) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeID", "body", m.VolumeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this snapshot v1 based on context it is used
func (m *SnapshotV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotV1) UnmarshalBinary(b []byte) error {
	var res SnapshotV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
