// Code generated by go-swagger; DO NOT EDIT.

package tenants_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenants config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTenantConfig(params *CreateTenantConfigParams, opts ...ClientOption) (*CreateTenantConfigOK, error)

	DeleteTenantConfig(params *DeleteTenantConfigParams, opts ...ClientOption) (*DeleteTenantConfigOK, error)

	GetTenantConfig(params *GetTenantConfigParams, opts ...ClientOption) (*GetTenantConfigOK, error)

	TackleHandler(params *TackleHandlerParams, opts ...ClientOption) (*TackleHandlerOK, error)

	UpdateTenantConfig(params *UpdateTenantConfigParams, opts ...ClientOption) (*UpdateTenantConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTenantConfig creates tenant config for the authenticated user
*/
func (a *Client) CreateTenantConfig(params *CreateTenantConfigParams, opts ...ClientOption) (*CreateTenantConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTenantConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTenantConfig",
		Method:             "POST",
		PathPattern:        "/v1/tenants-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTenantConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTenantConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateTenantConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTenantConfig deletes tenant config
*/
func (a *Client) DeleteTenantConfig(params *DeleteTenantConfigParams, opts ...ClientOption) (*DeleteTenantConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTenantConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTenantConfig",
		Method:             "DELETE",
		PathPattern:        "/v1/tenants-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTenantConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTenantConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteTenantConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTenantConfig gets tenant config for the authenticated user
*/
func (a *Client) GetTenantConfig(params *GetTenantConfigParams, opts ...ClientOption) (*GetTenantConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTenantConfig",
		Method:             "GET",
		PathPattern:        "/v1/tenants-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTenantConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTenantConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TackleHandler adapts tackle request into a config manager request
*/
func (a *Client) TackleHandler(params *TackleHandlerParams, opts ...ClientOption) (*TackleHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTackleHandlerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TackleHandler",
		Method:             "POST",
		PathPattern:        "/v1/tenants-config/tackle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TackleHandlerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TackleHandlerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TackleHandler: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTenantConfig creates or update tenant config
*/
func (a *Client) UpdateTenantConfig(params *UpdateTenantConfigParams, opts ...ClientOption) (*UpdateTenantConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTenantConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTenantConfig",
		Method:             "PUT",
		PathPattern:        "/v1/tenants-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTenantConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTenantConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateTenantConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
