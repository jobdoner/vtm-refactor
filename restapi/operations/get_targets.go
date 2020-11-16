// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTargetsHandlerFunc turns a function with the right signature into a get targets handler
type GetTargetsHandlerFunc func(GetTargetsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTargetsHandlerFunc) Handle(params GetTargetsParams) middleware.Responder {
	return fn(params)
}

// GetTargetsHandler interface for that can handle valid get targets params
type GetTargetsHandler interface {
	Handle(GetTargetsParams) middleware.Responder
}

// NewGetTargets creates a new http.Handler for the get targets operation
func NewGetTargets(ctx *middleware.Context, handler GetTargetsHandler) *GetTargets {
	return &GetTargets{Context: ctx, Handler: handler}
}

/*GetTargets swagger:route GET /export/targets getTargets

GetTargets get targets API

*/
type GetTargets struct {
	Context *middleware.Context
	Handler GetTargetsHandler
}

func (o *GetTargets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTargetsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}