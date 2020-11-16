// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MatchFoldersHandlerFunc turns a function with the right signature into a match folders handler
type MatchFoldersHandlerFunc func(MatchFoldersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MatchFoldersHandlerFunc) Handle(params MatchFoldersParams) middleware.Responder {
	return fn(params)
}

// MatchFoldersHandler interface for that can handle valid match folders params
type MatchFoldersHandler interface {
	Handle(MatchFoldersParams) middleware.Responder
}

// NewMatchFolders creates a new http.Handler for the match folders operation
func NewMatchFolders(ctx *middleware.Context, handler MatchFoldersHandler) *MatchFolders {
	return &MatchFolders{Context: ctx, Handler: handler}
}

/*MatchFolders swagger:route POST /folders/match matchFolders

MatchFolders match folders API

*/
type MatchFolders struct {
	Context *middleware.Context
	Handler MatchFoldersHandler
}

func (o *MatchFolders) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMatchFoldersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}