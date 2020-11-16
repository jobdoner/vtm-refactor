// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShowFolderResp show folder resp
//
// swagger:model ShowFolderResp
type ShowFolderResp struct {

	// data
	Data *ShowFolderRespData `json:"data,omitempty"`
}

// Validate validates this show folder resp
func (m *ShowFolderResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShowFolderResp) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShowFolderResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShowFolderResp) UnmarshalBinary(b []byte) error {
	var res ShowFolderResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ShowFolderRespData show folder resp data
//
// swagger:model ShowFolderRespData
type ShowFolderRespData struct {

	// b l s split
	BLSSplit bool `json:"BLSSplit,omitempty"`

	// campaign ID
	CampaignID int64 `json:"campaignID,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// frequency
	Frequency string `json:"frequency,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rules mask
	RulesMask string `json:"rulesMask,omitempty"`

	// tech ad group ID
	TechAdGroupID int64 `json:"techAdGroupID,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this show folder resp data
func (m *ShowFolderRespData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShowFolderRespData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShowFolderRespData) UnmarshalBinary(b []byte) error {
	var res ShowFolderRespData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
