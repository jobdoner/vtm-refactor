// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jobdoner/vtm-refactor/models"
)

// DeleteCampaignIDOKCode is the HTTP code returned for type DeleteCampaignIDOK
const DeleteCampaignIDOKCode int = 200

/*DeleteCampaignIDOK delete campaign

swagger:response deleteCampaignIdOK
*/
type DeleteCampaignIDOK struct {

	/*
	  In: Body
	*/
	Payload models.ErrorMess `json:"body,omitempty"`
}

// NewDeleteCampaignIDOK creates DeleteCampaignIDOK with default headers values
func NewDeleteCampaignIDOK() *DeleteCampaignIDOK {

	return &DeleteCampaignIDOK{}
}

// WithPayload adds the payload to the delete campaign Id o k response
func (o *DeleteCampaignIDOK) WithPayload(payload models.ErrorMess) *DeleteCampaignIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete campaign Id o k response
func (o *DeleteCampaignIDOK) SetPayload(payload models.ErrorMess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCampaignIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteCampaignIDNoContentCode is the HTTP code returned for type DeleteCampaignIDNoContent
const DeleteCampaignIDNoContentCode int = 204

/*DeleteCampaignIDNoContent No Content

swagger:response deleteCampaignIdNoContent
*/
type DeleteCampaignIDNoContent struct {
}

// NewDeleteCampaignIDNoContent creates DeleteCampaignIDNoContent with default headers values
func NewDeleteCampaignIDNoContent() *DeleteCampaignIDNoContent {

	return &DeleteCampaignIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteCampaignIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteCampaignIDInternalServerErrorCode is the HTTP code returned for type DeleteCampaignIDInternalServerError
const DeleteCampaignIDInternalServerErrorCode int = 500

/*DeleteCampaignIDInternalServerError Internal Server Error

swagger:response deleteCampaignIdInternalServerError
*/
type DeleteCampaignIDInternalServerError struct {
}

// NewDeleteCampaignIDInternalServerError creates DeleteCampaignIDInternalServerError with default headers values
func NewDeleteCampaignIDInternalServerError() *DeleteCampaignIDInternalServerError {

	return &DeleteCampaignIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteCampaignIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
