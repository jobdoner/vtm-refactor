// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewVtmRefactorAPI creates a new VtmRefactor instance
func NewVtmRefactorAPI(spec *loads.Document) *VtmRefactorAPI {
	return &VtmRefactorAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		BinProducer:  runtime.ByteStreamProducer(),
		JSONProducer: runtime.JSONProducer(),

		DeleteCampaignIDHandler: DeleteCampaignIDHandlerFunc(func(params DeleteCampaignIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteCampaignID has not yet been implemented")
		}),
		DeleteFoldersIDHandler: DeleteFoldersIDHandlerFunc(func(params DeleteFoldersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteFoldersID has not yet been implemented")
		}),
		DeleteFoldersMatchIDHandler: DeleteFoldersMatchIDHandlerFunc(func(params DeleteFoldersMatchIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteFoldersMatchID has not yet been implemented")
		}),
		GetAdgroupHandler: GetAdgroupHandlerFunc(func(params GetAdgroupParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAdgroup has not yet been implemented")
		}),
		GetCampaignHandler: GetCampaignHandlerFunc(func(params GetCampaignParams) middleware.Responder {
			return middleware.NotImplemented("operation GetCampaign has not yet been implemented")
		}),
		GetCampaignIDHandler: GetCampaignIDHandlerFunc(func(params GetCampaignIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetCampaignID has not yet been implemented")
		}),
		GetFoldersHandler: GetFoldersHandlerFunc(func(params GetFoldersParams) middleware.Responder {
			return middleware.NotImplemented("operation GetFolders has not yet been implemented")
		}),
		GetFoldersIDHandler: GetFoldersIDHandlerFunc(func(params GetFoldersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetFoldersID has not yet been implemented")
		}),
		GetSettingsObjectHandler: GetSettingsObjectHandlerFunc(func(params GetSettingsObjectParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSettingsObject has not yet been implemented")
		}),
		GetSettingsTechHandler: GetSettingsTechHandlerFunc(func(params GetSettingsTechParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSettingsTech has not yet been implemented")
		}),
		ListVideosHandler: ListVideosHandlerFunc(func(params ListVideosParams) middleware.Responder {
			return middleware.NotImplemented("operation ListVideos has not yet been implemented")
		}),
		PatchCampaignIDHandler: PatchCampaignIDHandlerFunc(func(params PatchCampaignIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PatchCampaignID has not yet been implemented")
		}),
		PatchFoldersIDHandler: PatchFoldersIDHandlerFunc(func(params PatchFoldersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PatchFoldersID has not yet been implemented")
		}),
		PostImportTargetsHandler: PostImportTargetsHandlerFunc(func(params PostImportTargetsParams) middleware.Responder {
			return middleware.NotImplemented("operation PostImportTargets has not yet been implemented")
		}),
		PostVideoHandler: PostVideoHandlerFunc(func(params PostVideoParams) middleware.Responder {
			return middleware.NotImplemented("operation PostVideo has not yet been implemented")
		}),
		CreateAdHandler: CreateAdHandlerFunc(func(params CreateAdParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateAd has not yet been implemented")
		}),
		CreateAdGroupHandler: CreateAdGroupHandlerFunc(func(params CreateAdGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateAdGroup has not yet been implemented")
		}),
		CreateCampaignHandler: CreateCampaignHandlerFunc(func(params CreateCampaignParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateCampaign has not yet been implemented")
		}),
		CreateFoldersHandler: CreateFoldersHandlerFunc(func(params CreateFoldersParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateFolders has not yet been implemented")
		}),
		CreateTechAdGroupHandler: CreateTechAdGroupHandlerFunc(func(params CreateTechAdGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateTechAdGroup has not yet been implemented")
		}),
		DeleteAdgroupHandler: DeleteAdgroupHandlerFunc(func(params DeleteAdgroupParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteAdgroup has not yet been implemented")
		}),
		GetAdHandler: GetAdHandlerFunc(func(params GetAdParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAd has not yet been implemented")
		}),
		GetCSVHandler: GetCSVHandlerFunc(func(params GetCSVParams) middleware.Responder {
			return middleware.NotImplemented("operation GetCSV has not yet been implemented")
		}),
		GetSettingsHandler: GetSettingsHandlerFunc(func(params GetSettingsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSettings has not yet been implemented")
		}),
		GetTargetsHandler: GetTargetsHandlerFunc(func(params GetTargetsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetTargets has not yet been implemented")
		}),
		ListMatchingFoldersHandler: ListMatchingFoldersHandlerFunc(func(params ListMatchingFoldersParams) middleware.Responder {
			return middleware.NotImplemented("operation ListMatchingFolders has not yet been implemented")
		}),
		MatchFoldersHandler: MatchFoldersHandlerFunc(func(params MatchFoldersParams) middleware.Responder {
			return middleware.NotImplemented("operation MatchFolders has not yet been implemented")
		}),
		ShowAdgroupHandler: ShowAdgroupHandlerFunc(func(params ShowAdgroupParams) middleware.Responder {
			return middleware.NotImplemented("operation ShowAdgroup has not yet been implemented")
		}),
		UploadVideoHandler: UploadVideoHandlerFunc(func(params UploadVideoParams) middleware.Responder {
			return middleware.NotImplemented("operation UploadVideo has not yet been implemented")
		}),
	}
}

/*VtmRefactorAPI vtm app */
type VtmRefactorAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// DeleteCampaignIDHandler sets the operation handler for the delete campaign ID operation
	DeleteCampaignIDHandler DeleteCampaignIDHandler
	// DeleteFoldersIDHandler sets the operation handler for the delete folders ID operation
	DeleteFoldersIDHandler DeleteFoldersIDHandler
	// DeleteFoldersMatchIDHandler sets the operation handler for the delete folders match ID operation
	DeleteFoldersMatchIDHandler DeleteFoldersMatchIDHandler
	// GetAdgroupHandler sets the operation handler for the get adgroup operation
	GetAdgroupHandler GetAdgroupHandler
	// GetCampaignHandler sets the operation handler for the get campaign operation
	GetCampaignHandler GetCampaignHandler
	// GetCampaignIDHandler sets the operation handler for the get campaign ID operation
	GetCampaignIDHandler GetCampaignIDHandler
	// GetFoldersHandler sets the operation handler for the get folders operation
	GetFoldersHandler GetFoldersHandler
	// GetFoldersIDHandler sets the operation handler for the get folders ID operation
	GetFoldersIDHandler GetFoldersIDHandler
	// GetSettingsObjectHandler sets the operation handler for the get settings object operation
	GetSettingsObjectHandler GetSettingsObjectHandler
	// GetSettingsTechHandler sets the operation handler for the get settings tech operation
	GetSettingsTechHandler GetSettingsTechHandler
	// ListVideosHandler sets the operation handler for the list videos operation
	ListVideosHandler ListVideosHandler
	// PatchCampaignIDHandler sets the operation handler for the patch campaign ID operation
	PatchCampaignIDHandler PatchCampaignIDHandler
	// PatchFoldersIDHandler sets the operation handler for the patch folders ID operation
	PatchFoldersIDHandler PatchFoldersIDHandler
	// PostImportTargetsHandler sets the operation handler for the post import targets operation
	PostImportTargetsHandler PostImportTargetsHandler
	// PostVideoHandler sets the operation handler for the post video operation
	PostVideoHandler PostVideoHandler
	// CreateAdHandler sets the operation handler for the create ad operation
	CreateAdHandler CreateAdHandler
	// CreateAdGroupHandler sets the operation handler for the create ad group operation
	CreateAdGroupHandler CreateAdGroupHandler
	// CreateCampaignHandler sets the operation handler for the create campaign operation
	CreateCampaignHandler CreateCampaignHandler
	// CreateFoldersHandler sets the operation handler for the create folders operation
	CreateFoldersHandler CreateFoldersHandler
	// CreateTechAdGroupHandler sets the operation handler for the create tech ad group operation
	CreateTechAdGroupHandler CreateTechAdGroupHandler
	// DeleteAdgroupHandler sets the operation handler for the delete adgroup operation
	DeleteAdgroupHandler DeleteAdgroupHandler
	// GetAdHandler sets the operation handler for the get ad operation
	GetAdHandler GetAdHandler
	// GetCSVHandler sets the operation handler for the get c s v operation
	GetCSVHandler GetCSVHandler
	// GetSettingsHandler sets the operation handler for the get settings operation
	GetSettingsHandler GetSettingsHandler
	// GetTargetsHandler sets the operation handler for the get targets operation
	GetTargetsHandler GetTargetsHandler
	// ListMatchingFoldersHandler sets the operation handler for the list matching folders operation
	ListMatchingFoldersHandler ListMatchingFoldersHandler
	// MatchFoldersHandler sets the operation handler for the match folders operation
	MatchFoldersHandler MatchFoldersHandler
	// ShowAdgroupHandler sets the operation handler for the show adgroup operation
	ShowAdgroupHandler ShowAdgroupHandler
	// UploadVideoHandler sets the operation handler for the upload video operation
	UploadVideoHandler UploadVideoHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *VtmRefactorAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *VtmRefactorAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *VtmRefactorAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *VtmRefactorAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *VtmRefactorAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *VtmRefactorAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *VtmRefactorAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the VtmRefactorAPI
func (o *VtmRefactorAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DeleteCampaignIDHandler == nil {
		unregistered = append(unregistered, "DeleteCampaignIDHandler")
	}
	if o.DeleteFoldersIDHandler == nil {
		unregistered = append(unregistered, "DeleteFoldersIDHandler")
	}
	if o.DeleteFoldersMatchIDHandler == nil {
		unregistered = append(unregistered, "DeleteFoldersMatchIDHandler")
	}
	if o.GetAdgroupHandler == nil {
		unregistered = append(unregistered, "GetAdgroupHandler")
	}
	if o.GetCampaignHandler == nil {
		unregistered = append(unregistered, "GetCampaignHandler")
	}
	if o.GetCampaignIDHandler == nil {
		unregistered = append(unregistered, "GetCampaignIDHandler")
	}
	if o.GetFoldersHandler == nil {
		unregistered = append(unregistered, "GetFoldersHandler")
	}
	if o.GetFoldersIDHandler == nil {
		unregistered = append(unregistered, "GetFoldersIDHandler")
	}
	if o.GetSettingsObjectHandler == nil {
		unregistered = append(unregistered, "GetSettingsObjectHandler")
	}
	if o.GetSettingsTechHandler == nil {
		unregistered = append(unregistered, "GetSettingsTechHandler")
	}
	if o.ListVideosHandler == nil {
		unregistered = append(unregistered, "ListVideosHandler")
	}
	if o.PatchCampaignIDHandler == nil {
		unregistered = append(unregistered, "PatchCampaignIDHandler")
	}
	if o.PatchFoldersIDHandler == nil {
		unregistered = append(unregistered, "PatchFoldersIDHandler")
	}
	if o.PostImportTargetsHandler == nil {
		unregistered = append(unregistered, "PostImportTargetsHandler")
	}
	if o.PostVideoHandler == nil {
		unregistered = append(unregistered, "PostVideoHandler")
	}
	if o.CreateAdHandler == nil {
		unregistered = append(unregistered, "CreateAdHandler")
	}
	if o.CreateAdGroupHandler == nil {
		unregistered = append(unregistered, "CreateAdGroupHandler")
	}
	if o.CreateCampaignHandler == nil {
		unregistered = append(unregistered, "CreateCampaignHandler")
	}
	if o.CreateFoldersHandler == nil {
		unregistered = append(unregistered, "CreateFoldersHandler")
	}
	if o.CreateTechAdGroupHandler == nil {
		unregistered = append(unregistered, "CreateTechAdGroupHandler")
	}
	if o.DeleteAdgroupHandler == nil {
		unregistered = append(unregistered, "DeleteAdgroupHandler")
	}
	if o.GetAdHandler == nil {
		unregistered = append(unregistered, "GetAdHandler")
	}
	if o.GetCSVHandler == nil {
		unregistered = append(unregistered, "GetCSVHandler")
	}
	if o.GetSettingsHandler == nil {
		unregistered = append(unregistered, "GetSettingsHandler")
	}
	if o.GetTargetsHandler == nil {
		unregistered = append(unregistered, "GetTargetsHandler")
	}
	if o.ListMatchingFoldersHandler == nil {
		unregistered = append(unregistered, "ListMatchingFoldersHandler")
	}
	if o.MatchFoldersHandler == nil {
		unregistered = append(unregistered, "MatchFoldersHandler")
	}
	if o.ShowAdgroupHandler == nil {
		unregistered = append(unregistered, "ShowAdgroupHandler")
	}
	if o.UploadVideoHandler == nil {
		unregistered = append(unregistered, "UploadVideoHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *VtmRefactorAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *VtmRefactorAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *VtmRefactorAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *VtmRefactorAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *VtmRefactorAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *VtmRefactorAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the vtm refactor API
func (o *VtmRefactorAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *VtmRefactorAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/campaign/{id}"] = NewDeleteCampaignID(o.context, o.DeleteCampaignIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/folders/{id}"] = NewDeleteFoldersID(o.context, o.DeleteFoldersIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/folders/match/{id}"] = NewDeleteFoldersMatchID(o.context, o.DeleteFoldersMatchIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/adgroup"] = NewGetAdgroup(o.context, o.GetAdgroupHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/campaign"] = NewGetCampaign(o.context, o.GetCampaignHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/campaign/{id}"] = NewGetCampaignID(o.context, o.GetCampaignIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/folders"] = NewGetFolders(o.context, o.GetFoldersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/folders/{id}"] = NewGetFoldersID(o.context, o.GetFoldersIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/settings/object"] = NewGetSettingsObject(o.context, o.GetSettingsObjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/settings/tech"] = NewGetSettingsTech(o.context, o.GetSettingsTechHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/video/folders/{folderID}"] = NewListVideos(o.context, o.ListVideosHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/campaign/{id}"] = NewPatchCampaignID(o.context, o.PatchCampaignIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/folders/{id}"] = NewPatchFoldersID(o.context, o.PatchFoldersIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/import/targets"] = NewPostImportTargets(o.context, o.PostImportTargetsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/video"] = NewPostVideo(o.context, o.PostVideoHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/ad"] = NewCreateAd(o.context, o.CreateAdHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/adgroup"] = NewCreateAdGroup(o.context, o.CreateAdGroupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/campaign"] = NewCreateCampaign(o.context, o.CreateCampaignHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/folders"] = NewCreateFolders(o.context, o.CreateFoldersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/adgroup/tech"] = NewCreateTechAdGroup(o.context, o.CreateTechAdGroupHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/adgroup/{id}"] = NewDeleteAdgroup(o.context, o.DeleteAdgroupHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ad"] = NewGetAd(o.context, o.GetAdHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/csv"] = NewGetCSV(o.context, o.GetCSVHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/settings/{list_name}"] = NewGetSettings(o.context, o.GetSettingsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/export/targets"] = NewGetTargets(o.context, o.GetTargetsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/folders/match"] = NewListMatchingFolders(o.context, o.ListMatchingFoldersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/folders/match"] = NewMatchFolders(o.context, o.MatchFoldersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/adgroup/{id}"] = NewShowAdgroup(o.context, o.ShowAdgroupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/video/upload/{folderID}"] = NewUploadVideo(o.context, o.UploadVideoHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *VtmRefactorAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *VtmRefactorAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *VtmRefactorAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *VtmRefactorAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *VtmRefactorAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
