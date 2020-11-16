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

// CreateAdGroupResp create ad group resp
//
// swagger:model CreateAdGroupResp
type CreateAdGroupResp struct {

	// data
	Data []*CreateAdGroupRespDataItems0 `json:"data"`
}

// Validate validates this create ad group resp
func (m *CreateAdGroupResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAdGroupResp) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAdGroupResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAdGroupResp) UnmarshalBinary(b []byte) error {
	var res CreateAdGroupResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateAdGroupRespDataItems0 create ad group resp data items0
//
// swagger:model CreateAdGroupRespDataItems0
type CreateAdGroupRespDataItems0 struct {

	// ad group ID
	AdGroupID []string `json:"adGroupID"`

	// ad schedule
	AdSchedule []string `json:"adSchedule"`

	// age
	Age []string `json:"age"`

	// audience
	Audience []string `json:"audience"`

	// campaign ID
	CampaignID int64 `json:"campaignID,omitempty"`

	// channel ID
	ChannelID []string `json:"channelID"`

	// custom audience
	CustomAudience []string `json:"customAudience"`

	// devices
	Devices []string `json:"devices"`

	// folder ID
	FolderID int64 `json:"folderID,omitempty"`

	// frequency
	Frequency []string `json:"frequency"`

	// gender
	Gender []string `json:"gender"`

	// id
	ID int64 `json:"id,omitempty"`

	// interest
	Interest []string `json:"interest"`

	// keyword
	Keyword []string `json:"keyword"`

	// languages
	Languages []string `json:"languages"`

	// location
	Location []string `json:"location"`

	// placement
	Placement []string `json:"placement"`

	// tech ad group ID
	TechAdGroupID int64 `json:"techAdGroupID,omitempty"`

	// topics
	Topics []string `json:"topics"`
}

// Validate validates this create ad group resp data items0
func (m *CreateAdGroupRespDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAdGroupRespDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAdGroupRespDataItems0) UnmarshalBinary(b []byte) error {
	var res CreateAdGroupRespDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}