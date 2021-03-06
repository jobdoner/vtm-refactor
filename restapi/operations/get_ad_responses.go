// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jobdoner/vtm-refactor/models"
)

// GetAdOKCode is the HTTP code returned for type GetAdOK
const GetAdOKCode int = 200

/*GetAdOK create campaign

swagger:response getAdOK
*/
type GetAdOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetAdResp `json:"body,omitempty"`
}

// NewGetAdOK creates GetAdOK with default headers values
func NewGetAdOK() *GetAdOK {

	return &GetAdOK{}
}

// WithPayload adds the payload to the get ad o k response
func (o *GetAdOK) WithPayload(payload *models.GetAdResp) *GetAdOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ad o k response
func (o *GetAdOK) SetPayload(payload *models.GetAdResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAdOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAdNoContentCode is the HTTP code returned for type GetAdNoContent
const GetAdNoContentCode int = 204

/*GetAdNoContent No Content

swagger:response getAdNoContent
*/
type GetAdNoContent struct {
}

// NewGetAdNoContent creates GetAdNoContent with default headers values
func NewGetAdNoContent() *GetAdNoContent {

	return &GetAdNoContent{}
}

// WriteResponse to the client
func (o *GetAdNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetAdInternalServerErrorCode is the HTTP code returned for type GetAdInternalServerError
const GetAdInternalServerErrorCode int = 500

/*GetAdInternalServerError Internal Server Error

swagger:response getAdInternalServerError
*/
type GetAdInternalServerError struct {
}

// NewGetAdInternalServerError creates GetAdInternalServerError with default headers values
func NewGetAdInternalServerError() *GetAdInternalServerError {

	return &GetAdInternalServerError{}
}

// WriteResponse to the client
func (o *GetAdInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
