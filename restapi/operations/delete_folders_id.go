// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteFoldersIDHandlerFunc turns a function with the right signature into a delete folders ID handler
type DeleteFoldersIDHandlerFunc func(DeleteFoldersIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFoldersIDHandlerFunc) Handle(params DeleteFoldersIDParams) middleware.Responder {
	return fn(params)
}

// DeleteFoldersIDHandler interface for that can handle valid delete folders ID params
type DeleteFoldersIDHandler interface {
	Handle(DeleteFoldersIDParams) middleware.Responder
}

// NewDeleteFoldersID creates a new http.Handler for the delete folders ID operation
func NewDeleteFoldersID(ctx *middleware.Context, handler DeleteFoldersIDHandler) *DeleteFoldersID {
	return &DeleteFoldersID{Context: ctx, Handler: handler}
}

/*DeleteFoldersID swagger:route DELETE /folders/{id} deleteFoldersId

DeleteFoldersID delete folders ID API

*/
type DeleteFoldersID struct {
	Context *middleware.Context
	Handler DeleteFoldersIDHandler
}

func (o *DeleteFoldersID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFoldersIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}