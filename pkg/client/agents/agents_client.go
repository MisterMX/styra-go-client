// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new agents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAgentInfo(params *DeleteAgentInfoParams, opts ...ClientOption) (*DeleteAgentInfoOK, error)

	GetAgentStatuses(params *GetAgentStatusesParams, opts ...ClientOption) (*GetAgentStatusesOK, error)

	PostAgentStatus(params *PostAgentStatusParams, opts ...ClientOption) (*PostAgentStatusOK, error)

	UpdateAgentStatus(params *UpdateAgentStatusParams, opts ...ClientOption) (*UpdateAgentStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAgentInfo deletes agent information
*/
func (a *Client) DeleteAgentInfo(params *DeleteAgentInfoParams, opts ...ClientOption) (*DeleteAgentInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAgentInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAgentInfo",
		Method:             "DELETE",
		PathPattern:        "/v1/agents/{kind}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAgentInfoReader{formats: a.formats},
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
	success, ok := result.(*DeleteAgentInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAgentInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAgentStatuses gets current agent statuses
*/
func (a *Client) GetAgentStatuses(params *GetAgentStatusesParams, opts ...ClientOption) (*GetAgentStatusesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentStatusesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAgentStatuses",
		Method:             "GET",
		PathPattern:        "/v1/agents/{kind}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAgentStatusesReader{formats: a.formats},
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
	success, ok := result.(*GetAgentStatusesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAgentStatuses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAgentStatus posts agent status
*/
func (a *Client) PostAgentStatus(params *PostAgentStatusParams, opts ...ClientOption) (*PostAgentStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAgentStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAgentStatus",
		Method:             "POST",
		PathPattern:        "/v1/agents/{kind}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAgentStatusReader{formats: a.formats},
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
	success, ok := result.(*PostAgentStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAgentStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAgentStatus updates agent status
*/
func (a *Client) UpdateAgentStatus(params *UpdateAgentStatusParams, opts ...ClientOption) (*UpdateAgentStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAgentStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateAgentStatus",
		Method:             "PUT",
		PathPattern:        "/v1/agents/{kind}/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAgentStatusReader{formats: a.formats},
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
	success, ok := result.(*UpdateAgentStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateAgentStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
