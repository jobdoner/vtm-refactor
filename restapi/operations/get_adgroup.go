// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAdgroupHandlerFunc turns a function with the right signature into a get adgroup handler
type GetAdgroupHandlerFunc func(GetAdgroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAdgroupHandlerFunc) Handle(params GetAdgroupParams) middleware.Responder {
	return fn(params)
}

// GetAdgroupHandler interface for that can handle valid get adgroup params
type GetAdgroupHandler interface {
	Handle(GetAdgroupParams) middleware.Responder
}

// NewGetAdgroup creates a new http.Handler for the get adgroup operation
func NewGetAdgroup(ctx *middleware.Context, handler GetAdgroupHandler) *GetAdgroup {
	return &GetAdgroup{Context: ctx, Handler: handler}
}

/*GetAdgroup swagger:route GET /adgroup getAdgroup

GetAdgroup get adgroup API

*/
type GetAdgroup struct {
	Context *middleware.Context
	Handler GetAdgroupHandler
}

func (o *GetAdgroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAdgroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
