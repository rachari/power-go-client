// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeSnapshotList volume snapshot list
//
// swagger:model VolumeSnapshotList
type VolumeSnapshotList struct {

	// The list of volume snapshots.
	VolumeSnapshots []*SnapshotV1 `json:"volumeSnapshots"`
}

// Validate validates this volume snapshot list
func (m *VolumeSnapshotList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeSnapshotList) validateVolumeSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeSnapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeSnapshots); i++ {
		if swag.IsZero(m.VolumeSnapshots[i]) { // not required
			continue
		}

		if m.VolumeSnapshots[i] != nil {
			if err := m.VolumeSnapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeSnapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeSnapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this volume snapshot list based on the context it is used
func (m *VolumeSnapshotList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumeSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeSnapshotList) contextValidateVolumeSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeSnapshots); i++ {

		if m.VolumeSnapshots[i] != nil {

			if swag.IsZero(m.VolumeSnapshots[i]) { // not required
				return nil
			}

			if err := m.VolumeSnapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeSnapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeSnapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeSnapshotList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeSnapshotList) UnmarshalBinary(b []byte) error {
	var res VolumeSnapshotList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
