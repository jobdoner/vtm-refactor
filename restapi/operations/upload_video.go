// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadVideoHandlerFunc turns a function with the right signature into a upload video handler
type UploadVideoHandlerFunc func(UploadVideoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadVideoHandlerFunc) Handle(params UploadVideoParams) middleware.Responder {
	return fn(params)
}

// UploadVideoHandler interface for that can handle valid upload video params
type UploadVideoHandler interface {
	Handle(UploadVideoParams) middleware.Responder
}

// NewUploadVideo creates a new http.Handler for the upload video operation
func NewUploadVideo(ctx *middleware.Context, handler UploadVideoHandler) *UploadVideo {
	return &UploadVideo{Context: ctx, Handler: handler}
}

/*UploadVideo swagger:route POST /video/upload/{folderID} uploadVideo

UploadVideo upload video API

*/
type UploadVideo struct {
	Context *middleware.Context
	Handler UploadVideoHandler
}

func (o *UploadVideo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUploadVideoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}