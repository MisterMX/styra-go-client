// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CommitFilesToSourceControlStack(params *CommitFilesToSourceControlStackParams, opts ...ClientOption) (*CommitFilesToSourceControlStackOK, error)

	CreateStack(params *CreateStackParams, opts ...ClientOption) (*CreateStackOK, error)

	DeleteStack(params *DeleteStackParams, opts ...ClientOption) (*DeleteStackOK, error)

	DeleteUserBranchStack(params *DeleteUserBranchStackParams, opts ...ClientOption) (*DeleteUserBranchStackOK, error)

	GetSourceControlFilesBranchStack(params *GetSourceControlFilesBranchStackParams, opts ...ClientOption) (*GetSourceControlFilesBranchStackOK, error)

	GetSourceControlFilesMasterStack(params *GetSourceControlFilesMasterStackParams, opts ...ClientOption) (*GetSourceControlFilesMasterStackOK, error)

	GetStack(params *GetStackParams, opts ...ClientOption) (*GetStackOK, error)

	ListStacks(params *ListStacksParams, opts ...ClientOption) (*ListStacksOK, error)

	SourceControlVerifyConfigStack(params *SourceControlVerifyConfigStackParams, opts ...ClientOption) (*SourceControlVerifyConfigStackOK, error)

	UpdateStack(params *UpdateStackParams, opts ...ClientOption) (*UpdateStackOK, error)

	ValidateStackCompliance(params *ValidateStackComplianceParams, opts ...ClientOption) (*ValidateStackComplianceOK, *ValidateStackComplianceAccepted, error)

	ValidateStackTests(params *ValidateStackTestsParams, opts ...ClientOption) (*ValidateStackTestsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CommitFilesToSourceControlStack commits files to stack source control

  Commit files to source control associated with a stack
*/
func (a *Client) CommitFilesToSourceControlStack(params *CommitFilesToSourceControlStackParams, opts ...ClientOption) (*CommitFilesToSourceControlStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitFilesToSourceControlStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CommitFilesToSourceControlStack",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/commits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommitFilesToSourceControlStackReader{formats: a.formats},
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
	success, ok := result.(*CommitFilesToSourceControlStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CommitFilesToSourceControlStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateStack creates a stack
*/
func (a *Client) CreateStack(params *CreateStackParams, opts ...ClientOption) (*CreateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateStack",
		Method:             "POST",
		PathPattern:        "/v1/stacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateStackReader{formats: a.formats},
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
	success, ok := result.(*CreateStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteStack deletes a stack
*/
func (a *Client) DeleteStack(params *DeleteStackParams, opts ...ClientOption) (*DeleteStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteStack",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/{stack}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStackReader{formats: a.formats},
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
	success, ok := result.(*DeleteStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUserBranchStack deletes a user owned branch
*/
func (a *Client) DeleteUserBranchStack(params *DeleteUserBranchStackParams, opts ...ClientOption) (*DeleteUserBranchStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserBranchStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserBranchStack",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserBranchStackReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserBranchStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUserBranchStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesBranchStack lists files in styra d a s created branch

  Gets the list of files for the branch that the Styra DAS creates when modifying rego in the Styra DAS UI and pushing the changes to GitHub in a branch for review.
*/
func (a *Client) GetSourceControlFilesBranchStack(params *GetSourceControlFilesBranchStackParams, opts ...ClientOption) (*GetSourceControlFilesBranchStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesBranchStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesBranchStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSourceControlFilesBranchStackReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesBranchStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesBranchStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesMasterStack lists files in current branch

  Gets the list of files in the currently chosen branch.
*/
func (a *Client) GetSourceControlFilesMasterStack(params *GetSourceControlFilesMasterStackParams, opts ...ClientOption) (*GetSourceControlFilesMasterStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesMasterStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesMasterStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/master",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSourceControlFilesMasterStackReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesMasterStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesMasterStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStack gets a stack configuration
*/
func (a *Client) GetStack(params *GetStackParams, opts ...ClientOption) (*GetStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStack",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{stack}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStackReader{formats: a.formats},
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
	success, ok := result.(*GetStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListStacks lists stacks
*/
func (a *Client) ListStacks(params *ListStacksParams, opts ...ClientOption) (*ListStacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListStacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListStacks",
		Method:             "GET",
		PathPattern:        "/v1/stacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListStacksReader{formats: a.formats},
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
	success, ok := result.(*ListStacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListStacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SourceControlVerifyConfigStack verifies git access

  Verifies that the repository can be accessed with the provided credentials
*/
func (a *Client) SourceControlVerifyConfigStack(params *SourceControlVerifyConfigStackParams, opts ...ClientOption) (*SourceControlVerifyConfigStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSourceControlVerifyConfigStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SourceControlVerifyConfigStack",
		Method:             "POST",
		PathPattern:        "/v1/stacks/source-control/verify-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SourceControlVerifyConfigStackReader{formats: a.formats},
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
	success, ok := result.(*SourceControlVerifyConfigStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SourceControlVerifyConfigStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateStack creates or update a stack
*/
func (a *Client) UpdateStack(params *UpdateStackParams, opts ...ClientOption) (*UpdateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateStack",
		Method:             "PUT",
		PathPattern:        "/v1/stacks/{stack}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateStackReader{formats: a.formats},
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
	success, ok := result.(*UpdateStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateStack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateStackCompliance validates stack compliance
*/
func (a *Client) ValidateStackCompliance(params *ValidateStackComplianceParams, opts ...ClientOption) (*ValidateStackComplianceOK, *ValidateStackComplianceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateStackComplianceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ValidateStackCompliance",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{stack}/validate/compliance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateStackComplianceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ValidateStackComplianceOK:
		return value, nil, nil
	case *ValidateStackComplianceAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateStackTests validates stack unit tests
*/
func (a *Client) ValidateStackTests(params *ValidateStackTestsParams, opts ...ClientOption) (*ValidateStackTestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateStackTestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ValidateStackTests",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{stack}/validate/tests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateStackTestsReader{formats: a.formats},
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
	success, ok := result.(*ValidateStackTestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ValidateStackTests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
