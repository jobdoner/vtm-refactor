// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetTargetsParams creates a new GetTargetsParams object
// no default values defined in spec.
func NewGetTargetsParams() GetTargetsParams {

	return GetTargetsParams{}
}

// GetTargetsParams contains all the bound params for the get targets operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTargets
type GetTargetsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	CampaignID *int64
	/*
	  In: query
	*/
	FolderID *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTargetsParams() beforehand.
func (o *GetTargetsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCampaignID, qhkCampaignID, _ := qs.GetOK("campaignID")
	if err := o.bindCampaignID(qCampaignID, qhkCampaignID, route.Formats); err != nil {
		res = append(res, err)
	}

	qFolderID, qhkFolderID, _ := qs.GetOK("folderID")
	if err := o.bindFolderID(qFolderID, qhkFolderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCampaignID binds and validates parameter CampaignID from query.
func (o *GetTargetsParams) bindCampaignID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("campaignID", "query", "int64", raw)
	}
	o.CampaignID = &value

	return nil
}

// bindFolderID binds and validates parameter FolderID from query.
func (o *GetTargetsParams) bindFolderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("folderID", "query", "int64", raw)
	}
	o.FolderID = &value

	return nil
}