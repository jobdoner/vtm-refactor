// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListVideosParams creates a new ListVideosParams object
// no default values defined in spec.
func NewListVideosParams() ListVideosParams {

	return ListVideosParams{}
}

// ListVideosParams contains all the bound params for the list videos operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListVideos
type ListVideosParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	FolderID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListVideosParams() beforehand.
func (o *ListVideosParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rFolderID, rhkFolderID, _ := route.Params.GetOK("folderID")
	if err := o.bindFolderID(rFolderID, rhkFolderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFolderID binds and validates parameter FolderID from path.
func (o *ListVideosParams) bindFolderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("folderID", "path", "int64", raw)
	}
	o.FolderID = value

	return nil
}
