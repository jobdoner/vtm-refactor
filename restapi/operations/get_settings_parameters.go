// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetSettingsParams creates a new GetSettingsParams object
// no default values defined in spec.
func NewGetSettingsParams() GetSettingsParams {

	return GetSettingsParams{}
}

// GetSettingsParams contains all the bound params for the get settings operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSettings
type GetSettingsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ListName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSettingsParams() beforehand.
func (o *GetSettingsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rListName, rhkListName, _ := route.Params.GetOK("list_name")
	if err := o.bindListName(rListName, rhkListName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindListName binds and validates parameter ListName from path.
func (o *GetSettingsParams) bindListName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ListName = raw

	return nil
}
