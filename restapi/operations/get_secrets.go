// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSecretsHandlerFunc turns a function with the right signature into a get secrets handler
type GetSecretsHandlerFunc func(GetSecretsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSecretsHandlerFunc) Handle(params GetSecretsParams) middleware.Responder {
	return fn(params)
}

// GetSecretsHandler interface for that can handle valid get secrets params
type GetSecretsHandler interface {
	Handle(GetSecretsParams) middleware.Responder
}

// NewGetSecrets creates a new http.Handler for the get secrets operation
func NewGetSecrets(ctx *middleware.Context, handler GetSecretsHandler) *GetSecrets {
	return &GetSecrets{Context: ctx, Handler: handler}
}

/*GetSecrets swagger:route GET /secrets getSecrets

get secrets configured in this sidecar

*/
type GetSecrets struct {
	Context *middleware.Context
	Handler GetSecretsHandler
}

func (o *GetSecrets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSecretsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
