// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new datasources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for datasources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDatasource(params *DeleteDatasourceParams, opts ...ClientOption) (*DeleteDatasourceOK, error)

	ExecuteDatasource(params *ExecuteDatasourceParams, opts ...ClientOption) (*ExecuteDatasourceOK, error)

	GetDatasource(params *GetDatasourceParams, opts ...ClientOption) (*GetDatasourceOK, error)

	ListDatasources(params *ListDatasourcesParams, opts ...ClientOption) (*ListDatasourcesOK, error)

	UpsertDatasource(params *UpsertDatasourceParams, opts ...ClientOption) (*UpsertDatasourceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteDatasource deletes a datasource
*/
func (a *Client) DeleteDatasource(params *DeleteDatasourceParams, opts ...ClientOption) (*DeleteDatasourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatasourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDatasource",
		Method:             "DELETE",
		PathPattern:        "/v1/datasources/{datasource}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDatasourceReader{formats: a.formats},
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
	success, ok := result.(*DeleteDatasourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteDatasource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExecuteDatasource executes a datasource
*/
func (a *Client) ExecuteDatasource(params *ExecuteDatasourceParams, opts ...ClientOption) (*ExecuteDatasourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteDatasourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecuteDatasource",
		Method:             "POST",
		PathPattern:        "/v1/datasources/{datasource}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteDatasourceReader{formats: a.formats},
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
	success, ok := result.(*ExecuteDatasourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecuteDatasource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDatasource gets a datasource
*/
func (a *Client) GetDatasource(params *GetDatasourceParams, opts ...ClientOption) (*GetDatasourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatasourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDatasource",
		Method:             "GET",
		PathPattern:        "/v1/datasources/{datasource}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatasourceReader{formats: a.formats},
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
	success, ok := result.(*GetDatasourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDatasource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDatasources lists datasources
*/
func (a *Client) ListDatasources(params *ListDatasourcesParams, opts ...ClientOption) (*ListDatasourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDatasourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListDatasources",
		Method:             "GET",
		PathPattern:        "/v1/datasources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDatasourcesReader{formats: a.formats},
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
	success, ok := result.(*ListDatasourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListDatasources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpsertDatasource upserts a datasource
*/
func (a *Client) UpsertDatasource(params *UpsertDatasourceParams, opts ...ClientOption) (*UpsertDatasourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertDatasourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpsertDatasource",
		Method:             "PUT",
		PathPattern:        "/v1/datasources/{datasource}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpsertDatasourceReader{formats: a.formats},
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
	success, ok := result.(*UpsertDatasourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpsertDatasource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
