// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteAdgroupOKCode is the HTTP code returned for type DeleteAdgroupOK
const DeleteAdgroupOKCode int = 200

/*DeleteAdgroupOK delete AdGroup

swagger:response deleteAdgroupOK
*/
type DeleteAdgroupOK struct {
}

// NewDeleteAdgroupOK creates DeleteAdgroupOK with default headers values
func NewDeleteAdgroupOK() *DeleteAdgroupOK {

	return &DeleteAdgroupOK{}
}

// WriteResponse to the client
func (o *DeleteAdgroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteAdgroupNoContentCode is the HTTP code returned for type DeleteAdgroupNoContent
const DeleteAdgroupNoContentCode int = 204

/*DeleteAdgroupNoContent No Content

swagger:response deleteAdgroupNoContent
*/
type DeleteAdgroupNoContent struct {
}

// NewDeleteAdgroupNoContent creates DeleteAdgroupNoContent with default headers values
func NewDeleteAdgroupNoContent() *DeleteAdgroupNoContent {

	return &DeleteAdgroupNoContent{}
}

// WriteResponse to the client
func (o *DeleteAdgroupNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteAdgroupInternalServerErrorCode is the HTTP code returned for type DeleteAdgroupInternalServerError
const DeleteAdgroupInternalServerErrorCode int = 500

/*DeleteAdgroupInternalServerError Internal Server Error

swagger:response deleteAdgroupInternalServerError
*/
type DeleteAdgroupInternalServerError struct {
}

// NewDeleteAdgroupInternalServerError creates DeleteAdgroupInternalServerError with default headers values
func NewDeleteAdgroupInternalServerError() *DeleteAdgroupInternalServerError {

	return &DeleteAdgroupInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteAdgroupInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
