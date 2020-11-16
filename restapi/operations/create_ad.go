// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAdHandlerFunc turns a function with the right signature into a create ad handler
type CreateAdHandlerFunc func(CreateAdParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAdHandlerFunc) Handle(params CreateAdParams) middleware.Responder {
	return fn(params)
}

// CreateAdHandler interface for that can handle valid create ad params
type CreateAdHandler interface {
	Handle(CreateAdParams) middleware.Responder
}

// NewCreateAd creates a new http.Handler for the create ad operation
func NewCreateAd(ctx *middleware.Context, handler CreateAdHandler) *CreateAd {
	return &CreateAd{Context: ctx, Handler: handler}
}

/*CreateAd swagger:route POST /ad createAd

CreateAd create ad API

*/
type CreateAd struct {
	Context *middleware.Context
	Handler CreateAdHandler
}

func (o *CreateAd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAdParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}