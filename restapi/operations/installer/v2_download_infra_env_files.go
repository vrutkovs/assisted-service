// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2DownloadInfraEnvFilesHandlerFunc turns a function with the right signature into a v2 download infra env files handler
type V2DownloadInfraEnvFilesHandlerFunc func(V2DownloadInfraEnvFilesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2DownloadInfraEnvFilesHandlerFunc) Handle(params V2DownloadInfraEnvFilesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2DownloadInfraEnvFilesHandler interface for that can handle valid v2 download infra env files params
type V2DownloadInfraEnvFilesHandler interface {
	Handle(V2DownloadInfraEnvFilesParams, interface{}) middleware.Responder
}

// NewV2DownloadInfraEnvFiles creates a new http.Handler for the v2 download infra env files operation
func NewV2DownloadInfraEnvFiles(ctx *middleware.Context, handler V2DownloadInfraEnvFilesHandler) *V2DownloadInfraEnvFiles {
	return &V2DownloadInfraEnvFiles{Context: ctx, Handler: handler}
}

/*V2DownloadInfraEnvFiles swagger:route GET /v2/infra-envs/{infra_env_id}/downloads/files installer v2DownloadInfraEnvFiles

Downloads the customized ignition file for this host

*/
type V2DownloadInfraEnvFiles struct {
	Context *middleware.Context
	Handler V2DownloadInfraEnvFilesHandler
}

func (o *V2DownloadInfraEnvFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV2DownloadInfraEnvFilesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
