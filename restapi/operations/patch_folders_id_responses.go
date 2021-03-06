// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jobdoner/vtm-refactor/models"
)

// PatchFoldersIDOKCode is the HTTP code returned for type PatchFoldersIDOK
const PatchFoldersIDOKCode int = 200

/*PatchFoldersIDOK update folder

swagger:response patchFoldersIdOK
*/
type PatchFoldersIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateFolderResp `json:"body,omitempty"`
}

// NewPatchFoldersIDOK creates PatchFoldersIDOK with default headers values
func NewPatchFoldersIDOK() *PatchFoldersIDOK {

	return &PatchFoldersIDOK{}
}

// WithPayload adds the payload to the patch folders Id o k response
func (o *PatchFoldersIDOK) WithPayload(payload *models.CreateFolderResp) *PatchFoldersIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch folders Id o k response
func (o *PatchFoldersIDOK) SetPayload(payload *models.CreateFolderResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchFoldersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchFoldersIDNoContentCode is the HTTP code returned for type PatchFoldersIDNoContent
const PatchFoldersIDNoContentCode int = 204

/*PatchFoldersIDNoContent No Content

swagger:response patchFoldersIdNoContent
*/
type PatchFoldersIDNoContent struct {
}

// NewPatchFoldersIDNoContent creates PatchFoldersIDNoContent with default headers values
func NewPatchFoldersIDNoContent() *PatchFoldersIDNoContent {

	return &PatchFoldersIDNoContent{}
}

// WriteResponse to the client
func (o *PatchFoldersIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PatchFoldersIDInternalServerErrorCode is the HTTP code returned for type PatchFoldersIDInternalServerError
const PatchFoldersIDInternalServerErrorCode int = 500

/*PatchFoldersIDInternalServerError Internal Server Error

swagger:response patchFoldersIdInternalServerError
*/
type PatchFoldersIDInternalServerError struct {
}

// NewPatchFoldersIDInternalServerError creates PatchFoldersIDInternalServerError with default headers values
func NewPatchFoldersIDInternalServerError() *PatchFoldersIDInternalServerError {

	return &PatchFoldersIDInternalServerError{}
}

// WriteResponse to the client
func (o *PatchFoldersIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
