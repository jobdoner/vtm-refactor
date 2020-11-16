// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCampaignHandlerFunc turns a function with the right signature into a get campaign handler
type GetCampaignHandlerFunc func(GetCampaignParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCampaignHandlerFunc) Handle(params GetCampaignParams) middleware.Responder {
	return fn(params)
}

// GetCampaignHandler interface for that can handle valid get campaign params
type GetCampaignHandler interface {
	Handle(GetCampaignParams) middleware.Responder
}

// NewGetCampaign creates a new http.Handler for the get campaign operation
func NewGetCampaign(ctx *middleware.Context, handler GetCampaignHandler) *GetCampaign {
	return &GetCampaign{Context: ctx, Handler: handler}
}

/*GetCampaign swagger:route GET /campaign getCampaign

GetCampaign get campaign API

*/
type GetCampaign struct {
	Context *middleware.Context
	Handler GetCampaignHandler
}

func (o *GetCampaign) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCampaignParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}