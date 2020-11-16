// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAdGroupHandlerFunc turns a function with the right signature into a create ad group handler
type CreateAdGroupHandlerFunc func(CreateAdGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAdGroupHandlerFunc) Handle(params CreateAdGroupParams) middleware.Responder {
	return fn(params)
}

// CreateAdGroupHandler interface for that can handle valid create ad group params
type CreateAdGroupHandler interface {
	Handle(CreateAdGroupParams) middleware.Responder
}

// NewCreateAdGroup creates a new http.Handler for the create ad group operation
func NewCreateAdGroup(ctx *middleware.Context, handler CreateAdGroupHandler) *CreateAdGroup {
	return &CreateAdGroup{Context: ctx, Handler: handler}
}

/*CreateAdGroup swagger:route POST /adgroup createAdGroup

CreateAdGroup create ad group API

*/
type CreateAdGroup struct {
	Context *middleware.Context
	Handler CreateAdGroupHandler
}

func (o *CreateAdGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAdGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
