// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListMatchingFoldersResp list matching folders resp
//
// swagger:model ListMatchingFoldersResp
type ListMatchingFoldersResp []*ListMatchingFoldersRespItems0

// Validate validates this list matching folders resp
func (m ListMatchingFoldersResp) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ListMatchingFoldersRespItems0 list matching folders resp items0
//
// swagger:model ListMatchingFoldersRespItems0
type ListMatchingFoldersRespItems0 struct {

	// audience folder ID
	AudienceFolderID int64 `json:"audienceFolderID,omitempty"`

	// audience folder name
	AudienceFolderName string `json:"audienceFolderName,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// video folder ID
	VideoFolderID int64 `json:"videoFolderID,omitempty"`

	// video folder name
	VideoFolderName string `json:"videoFolderName,omitempty"`
}

// Validate validates this list matching folders resp items0
func (m *ListMatchingFoldersRespItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListMatchingFoldersRespItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListMatchingFoldersRespItems0) UnmarshalBinary(b []byte) error {
	var res ListMatchingFoldersRespItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}