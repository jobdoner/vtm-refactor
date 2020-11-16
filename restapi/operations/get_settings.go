// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSettingsHandlerFunc turns a function with the right signature into a get settings handler
type GetSettingsHandlerFunc func(GetSettingsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSettingsHandlerFunc) Handle(params GetSettingsParams) middleware.Responder {
	return fn(params)
}

// GetSettingsHandler interface for that can handle valid get settings params
type GetSettingsHandler interface {
	Handle(GetSettingsParams) middleware.Responder
}

// NewGetSettings creates a new http.Handler for the get settings operation
func NewGetSettings(ctx *middleware.Context, handler GetSettingsHandler) *GetSettings {
	return &GetSettings{Context: ctx, Handler: handler}
}

/*GetSettings swagger:route GET /settings/{list_name} getSettings

GetSettings get settings API

*/
type GetSettings struct {
	Context *middleware.Context
	Handler GetSettingsHandler
}

func (o *GetSettings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSettingsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}