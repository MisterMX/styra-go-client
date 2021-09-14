// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workspace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workspace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CommitFilesToSourceControlWorkspace(params *CommitFilesToSourceControlWorkspaceParams, opts ...ClientOption) (*CommitFilesToSourceControlWorkspaceOK, error)

	DeleteUserBranchWorkspace(params *DeleteUserBranchWorkspaceParams, opts ...ClientOption) (*DeleteUserBranchWorkspaceOK, error)

	GetSourceControlFilesBranchWorkspace(params *GetSourceControlFilesBranchWorkspaceParams, opts ...ClientOption) (*GetSourceControlFilesBranchWorkspaceOK, error)

	GetSourceControlFilesMasterWorkspace(params *GetSourceControlFilesMasterWorkspaceParams, opts ...ClientOption) (*GetSourceControlFilesMasterWorkspaceOK, error)

	GetWorkspace(params *GetWorkspaceParams, opts ...ClientOption) (*GetWorkspaceOK, error)

	UpdateWorkspace(params *UpdateWorkspaceParams, opts ...ClientOption) (*UpdateWorkspaceOK, error)

	Func3(params *Func3Params, opts ...ClientOption) (*Func3OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CommitFilesToSourceControlWorkspace commits files to source control associated with a workspace
*/
func (a *Client) CommitFilesToSourceControlWorkspace(params *CommitFilesToSourceControlWorkspaceParams, opts ...ClientOption) (*CommitFilesToSourceControlWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitFilesToSourceControlWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CommitFilesToSourceControlWorkspace",
		Method:             "POST",
		PathPattern:        "/v1/workspace/{id}/commits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommitFilesToSourceControlWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*CommitFilesToSourceControlWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CommitFilesToSourceControlWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUserBranchWorkspace deletes a user owned branch
*/
func (a *Client) DeleteUserBranchWorkspace(params *DeleteUserBranchWorkspaceParams, opts ...ClientOption) (*DeleteUserBranchWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserBranchWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserBranchWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v1/workspace/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserBranchWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserBranchWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUserBranchWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesBranchWorkspace gets the list of files for the styra d a s created branch

  Gets the list of files for the branch that the Styra DAS creates when modifying rego in the Styra DAS UI and pushing the changes to GitHub in a branch for review.
*/
func (a *Client) GetSourceControlFilesBranchWorkspace(params *GetSourceControlFilesBranchWorkspaceParams, opts ...ClientOption) (*GetSourceControlFilesBranchWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesBranchWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesBranchWorkspace",
		Method:             "GET",
		PathPattern:        "/v1/workspace/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSourceControlFilesBranchWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesBranchWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesBranchWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesMasterWorkspace gets the list of files in the currently chosen branch
*/
func (a *Client) GetSourceControlFilesMasterWorkspace(params *GetSourceControlFilesMasterWorkspaceParams, opts ...ClientOption) (*GetSourceControlFilesMasterWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesMasterWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesMasterWorkspace",
		Method:             "GET",
		PathPattern:        "/v1/workspace/{id}/master",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSourceControlFilesMasterWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesMasterWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesMasterWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkspace gets workspace
*/
func (a *Client) GetWorkspace(params *GetWorkspaceParams, opts ...ClientOption) (*GetWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWorkspace",
		Method:             "GET",
		PathPattern:        "/v1/workspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*GetWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateWorkspace updates workspace
*/
func (a *Client) UpdateWorkspace(params *UpdateWorkspaceParams, opts ...ClientOption) (*UpdateWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateWorkspace",
		Method:             "PUT",
		PathPattern:        "/v1/workspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*UpdateWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Func3 verifies git configuration

  verifies that the repository can be accessed with the provided credentials
*/
func (a *Client) Func3(params *Func3Params, opts ...ClientOption) (*Func3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFunc3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "func3",
		Method:             "POST",
		PathPattern:        "/v1/workspace/source-control/verify-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Func3Reader{formats: a.formats},
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
	success, ok := result.(*Func3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for func3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
