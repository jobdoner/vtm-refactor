// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSettingsTechHandlerFunc turns a function with the right signature into a get settings tech handler
type GetSettingsTechHandlerFunc func(GetSettingsTechParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSettingsTechHandlerFunc) Handle(params GetSettingsTechParams) middleware.Responder {
	return fn(params)
}

// GetSettingsTechHandler interface for that can handle valid get settings tech params
type GetSettingsTechHandler interface {
	Handle(GetSettingsTechParams) middleware.Responder
}

// NewGetSettingsTech creates a new http.Handler for the get settings tech operation
func NewGetSettingsTech(ctx *middleware.Context, handler GetSettingsTechHandler) *GetSettingsTech {
	return &GetSettingsTech{Context: ctx, Handler: handler}
}

/*GetSettingsTech swagger:route GET /settings/tech getSettingsTech

GetSettingsTech get settings tech API

*/
type GetSettingsTech struct {
	Context *middleware.Context
	Handler GetSettingsTechHandler
}

func (o *GetSettingsTech) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSettingsTechParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
