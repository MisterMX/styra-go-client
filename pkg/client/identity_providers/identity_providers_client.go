// Code generated by go-swagger; DO NOT EDIT.

package identity_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new identity providers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for identity providers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProvider(params *CreateProviderParams, opts ...ClientOption) (*CreateProviderOK, error)

	DeleteProvider(params *DeleteProviderParams, opts ...ClientOption) (*DeleteProviderOK, error)

	GetProvider(params *GetProviderParams, opts ...ClientOption) (*GetProviderOK, error)

	ListProviders(params *ListProvidersParams, opts ...ClientOption) (*ListProvidersOK, error)

	UpdateProvider(params *UpdateProviderParams, opts ...ClientOption) (*UpdateProviderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateProvider creates provider
*/
func (a *Client) CreateProvider(params *CreateProviderParams, opts ...ClientOption) (*CreateProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateProvider",
		Method:             "POST",
		PathPattern:        "/v1/identity-providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProviderReader{formats: a.formats},
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
	success, ok := result.(*CreateProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProvider deletes provider
*/
func (a *Client) DeleteProvider(params *DeleteProviderParams, opts ...ClientOption) (*DeleteProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteProvider",
		Method:             "DELETE",
		PathPattern:        "/v1/identity-providers/{providerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProviderReader{formats: a.formats},
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
	success, ok := result.(*DeleteProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProvider gets provider
*/
func (a *Client) GetProvider(params *GetProviderParams, opts ...ClientOption) (*GetProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProvider",
		Method:             "GET",
		PathPattern:        "/v1/identity-providers/{providerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProviderReader{formats: a.formats},
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
	success, ok := result.(*GetProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListProviders lists providers
*/
func (a *Client) ListProviders(params *ListProvidersParams, opts ...ClientOption) (*ListProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProvidersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListProviders",
		Method:             "GET",
		PathPattern:        "/v1/identity-providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProvidersReader{formats: a.formats},
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
	success, ok := result.(*ListProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListProviders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProvider creates or uppdate provider
*/
func (a *Client) UpdateProvider(params *UpdateProviderParams, opts ...ClientOption) (*UpdateProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateProvider",
		Method:             "PUT",
		PathPattern:        "/v1/identity-providers/{providerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProviderReader{formats: a.formats},
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
	success, ok := result.(*UpdateProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
