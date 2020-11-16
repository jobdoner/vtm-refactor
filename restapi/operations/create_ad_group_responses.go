// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jobdoner/vtm-refactor/models"
)

// CreateAdGroupOKCode is the HTTP code returned for type CreateAdGroupOK
const CreateAdGroupOKCode int = 200

/*CreateAdGroupOK create adgroup

swagger:response createAdGroupOK
*/
type CreateAdGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateAdGroupResp `json:"body,omitempty"`
}

// NewCreateAdGroupOK creates CreateAdGroupOK with default headers values
func NewCreateAdGroupOK() *CreateAdGroupOK {

	return &CreateAdGroupOK{}
}

// WithPayload adds the payload to the create ad group o k response
func (o *CreateAdGroupOK) WithPayload(payload *models.CreateAdGroupResp) *CreateAdGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create ad group o k response
func (o *CreateAdGroupOK) SetPayload(payload *models.CreateAdGroupResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdGroupConflictCode is the HTTP code returned for type CreateAdGroupConflict
const CreateAdGroupConflictCode int = 409

/*CreateAdGroupConflict AdGroup already exis

swagger:response createAdGroupConflict
*/
type CreateAdGroupConflict struct {
}

// NewCreateAdGroupConflict creates CreateAdGroupConflict with default headers values
func NewCreateAdGroupConflict() *CreateAdGroupConflict {

	return &CreateAdGroupConflict{}
}

// WriteResponse to the client
func (o *CreateAdGroupConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
