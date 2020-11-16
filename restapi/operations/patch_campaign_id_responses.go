// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jobdoner/vtm-refactor/models"
)

// PatchCampaignIDOKCode is the HTTP code returned for type PatchCampaignIDOK
const PatchCampaignIDOKCode int = 200

/*PatchCampaignIDOK update campaign

swagger:response patchCampaignIdOK
*/
type PatchCampaignIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateCampaignResp `json:"body,omitempty"`
}

// NewPatchCampaignIDOK creates PatchCampaignIDOK with default headers values
func NewPatchCampaignIDOK() *PatchCampaignIDOK {

	return &PatchCampaignIDOK{}
}

// WithPayload adds the payload to the patch campaign Id o k response
func (o *PatchCampaignIDOK) WithPayload(payload *models.CreateCampaignResp) *PatchCampaignIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch campaign Id o k response
func (o *PatchCampaignIDOK) SetPayload(payload *models.CreateCampaignResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCampaignIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCampaignIDNoContentCode is the HTTP code returned for type PatchCampaignIDNoContent
const PatchCampaignIDNoContentCode int = 204

/*PatchCampaignIDNoContent No Content

swagger:response patchCampaignIdNoContent
*/
type PatchCampaignIDNoContent struct {
}

// NewPatchCampaignIDNoContent creates PatchCampaignIDNoContent with default headers values
func NewPatchCampaignIDNoContent() *PatchCampaignIDNoContent {

	return &PatchCampaignIDNoContent{}
}

// WriteResponse to the client
func (o *PatchCampaignIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PatchCampaignIDInternalServerErrorCode is the HTTP code returned for type PatchCampaignIDInternalServerError
const PatchCampaignIDInternalServerErrorCode int = 500

/*PatchCampaignIDInternalServerError Internal Server Error

swagger:response patchCampaignIdInternalServerError
*/
type PatchCampaignIDInternalServerError struct {
}

// NewPatchCampaignIDInternalServerError creates PatchCampaignIDInternalServerError with default headers values
func NewPatchCampaignIDInternalServerError() *PatchCampaignIDInternalServerError {

	return &PatchCampaignIDInternalServerError{}
}

// WriteResponse to the client
func (o *PatchCampaignIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
