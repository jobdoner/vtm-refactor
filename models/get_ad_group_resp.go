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

// GetAdGroupResp get ad group resp
//
// swagger:model GetAdGroupResp
type GetAdGroupResp struct {

	// data
	Data []*GetAdGroupRespDataItems0 `json:"data"`
}

// Validate validates this get ad group resp
func (m *GetAdGroupResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAdGroupResp) validateData(formats strfmt.Registry) error {

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
func (m *GetAdGroupResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAdGroupResp) UnmarshalBinary(b []byte) error {
	var res GetAdGroupResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetAdGroupRespDataItems0 get ad group resp data items0
//
// swagger:model GetAdGroupRespDataItems0
type GetAdGroupRespDataItems0 struct {

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

// Validate validates this get ad group resp data items0
func (m *GetAdGroupRespDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetAdGroupRespDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAdGroupRespDataItems0) UnmarshalBinary(b []byte) error {
	var res GetAdGroupRespDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}