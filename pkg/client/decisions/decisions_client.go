// Code generated by go-swagger; DO NOT EDIT.

package decisions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new decisions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for decisions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDecision(params *GetDecisionParams, opts ...ClientOption) (*GetDecisionOK, error)

	SearchDecisionLogsGet(params *SearchDecisionLogsGetParams, opts ...ClientOption) (*SearchDecisionLogsGetOK, error)

	SearchDecisionLogsPost(params *SearchDecisionLogsPostParams, opts ...ClientOption) (*SearchDecisionLogsPostOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDecision gets a single decision
*/
func (a *Client) GetDecision(params *GetDecisionParams, opts ...ClientOption) (*GetDecisionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDecisionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDecision",
		Method:             "GET",
		PathPattern:        "/v1/decisions/{cursor}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDecisionReader{formats: a.formats},
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
	success, ok := result.(*GetDecisionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDecision: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchDecisionLogsGet searches decision logs
*/
func (a *Client) SearchDecisionLogsGet(params *SearchDecisionLogsGetParams, opts ...ClientOption) (*SearchDecisionLogsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchDecisionLogsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchDecisionLogsGet",
		Method:             "GET",
		PathPattern:        "/v1/decisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchDecisionLogsGetReader{formats: a.formats},
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
	success, ok := result.(*SearchDecisionLogsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchDecisionLogsGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchDecisionLogsPost searches decision logs
*/
func (a *Client) SearchDecisionLogsPost(params *SearchDecisionLogsPostParams, opts ...ClientOption) (*SearchDecisionLogsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchDecisionLogsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchDecisionLogsPost",
		Method:             "POST",
		PathPattern:        "/v1/decisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchDecisionLogsPostReader{formats: a.formats},
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
	success, ok := result.(*SearchDecisionLogsPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchDecisionLogsPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
