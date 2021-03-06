// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAdgroupHandlerFunc turns a function with the right signature into a delete adgroup handler
type DeleteAdgroupHandlerFunc func(DeleteAdgroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAdgroupHandlerFunc) Handle(params DeleteAdgroupParams) middleware.Responder {
	return fn(params)
}

// DeleteAdgroupHandler interface for that can handle valid delete adgroup params
type DeleteAdgroupHandler interface {
	Handle(DeleteAdgroupParams) middleware.Responder
}

// NewDeleteAdgroup creates a new http.Handler for the delete adgroup operation
func NewDeleteAdgroup(ctx *middleware.Context, handler DeleteAdgroupHandler) *DeleteAdgroup {
	return &DeleteAdgroup{Context: ctx, Handler: handler}
}

/*DeleteAdgroup swagger:route DELETE /adgroup/{id} deleteAdgroup

DeleteAdgroup delete adgroup API

*/
type DeleteAdgroup struct {
	Context *middleware.Context
	Handler DeleteAdgroupHandler
}

func (o *DeleteAdgroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAdgroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
