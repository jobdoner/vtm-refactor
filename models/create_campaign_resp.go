// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateCampaignResp create campaign resp
//
// swagger:model CreateCampaignResp
type CreateCampaignResp struct {

	// data
	Data *CreateCampaignRespData `json:"data,omitempty"`
}

// Validate validates this create campaign resp
func (m *CreateCampaignResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCampaignResp) validateData(formats strfmt.Registry) error {

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
func (m *CreateCampaignResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCampaignResp) UnmarshalBinary(b []byte) error {
	var res CreateCampaignResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateCampaignRespData create campaign resp data
//
// swagger:model CreateCampaignRespData
type CreateCampaignRespData struct {

	// campaign ID
	CampaignID int64 `json:"campaignID,omitempty"`

	// tech ad group ID
	TechAdGroupID int64 `json:"techAdGroupID,omitempty"`
}

// Validate validates this create campaign resp data
func (m *CreateCampaignRespData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateCampaignRespData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCampaignRespData) UnmarshalBinary(b []byte) error {
	var res CreateCampaignRespData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
