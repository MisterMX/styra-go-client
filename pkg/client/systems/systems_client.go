// Code generated by go-swagger; DO NOT EDIT.

package systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new systems API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for systems API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CommitFilesToSourceControlSystem(params *CommitFilesToSourceControlSystemParams, opts ...ClientOption) (*CommitFilesToSourceControlSystemOK, error)

	CreateSystem(params *CreateSystemParams, opts ...ClientOption) (*CreateSystemOK, error)

	DeleteSystem(params *DeleteSystemParams, opts ...ClientOption) (*DeleteSystemOK, error)

	DeleteUserBranchSystem(params *DeleteUserBranchSystemParams, opts ...ClientOption) (*DeleteUserBranchSystemOK, error)

	GetAsset(params *GetAssetParams, writer io.Writer, opts ...ClientOption) (*GetAssetOK, error)

	GetDefaultPolicies(params *GetDefaultPoliciesParams, opts ...ClientOption) (*GetDefaultPoliciesOK, error)

	GetDefaultPolicy(params *GetDefaultPolicyParams, opts ...ClientOption) (*GetDefaultPolicyOK, error)

	GetInstructions(params *GetInstructionsParams, opts ...ClientOption) (*GetInstructionsOK, error)

	GetOPADiscoveryConfig(params *GetOPADiscoveryConfigParams, opts ...ClientOption) (*GetOPADiscoveryConfigOK, error)

	GetSourceControlFilesBranchSystem(params *GetSourceControlFilesBranchSystemParams, opts ...ClientOption) (*GetSourceControlFilesBranchSystemOK, error)

	GetSourceControlFilesMasterSystem(params *GetSourceControlFilesMasterSystemParams, opts ...ClientOption) (*GetSourceControlFilesMasterSystemOK, error)

	GetSystem(params *GetSystemParams, opts ...ClientOption) (*GetSystemOK, error)

	GetSystemAgents(params *GetSystemAgentsParams, opts ...ClientOption) (*GetSystemAgentsOK, error)

	GetSystemBundle(params *GetSystemBundleParams, opts ...ClientOption) (*GetSystemBundleOK, error)

	GetSystemBundleDeploy(params *GetSystemBundleDeployParams, opts ...ClientOption) (*GetSystemBundleDeployOK, error)

	GetSystemBundleDetails(params *GetSystemBundleDetailsParams, opts ...ClientOption) (*GetSystemBundleDetailsOK, error)

	GetSystemBundles(params *GetSystemBundlesParams, opts ...ClientOption) (*GetSystemBundlesOK, error)

	HandleSystemMetrics(params *HandleSystemMetricsParams, opts ...ClientOption) (*HandleSystemMetricsOK, error)

	ListSystems(params *ListSystemsParams, opts ...ClientOption) (*ListSystemsOK, error)

	RuleSuggestions(params *RuleSuggestionsParams, opts ...ClientOption) (*RuleSuggestionsOK, error)

	TranslateExternalIds(params *TranslateExternalIdsParams, opts ...ClientOption) (*TranslateExternalIdsOK, error)

	UpdateSystem(params *UpdateSystemParams, opts ...ClientOption) (*UpdateSystemOK, error)

	UpdateSystemBundleDeploy(params *UpdateSystemBundleDeployParams, opts ...ClientOption) (*UpdateSystemBundleDeployOK, error)

	ValidateSystemCompliance(params *ValidateSystemComplianceParams, opts ...ClientOption) (*ValidateSystemComplianceOK, *ValidateSystemComplianceAccepted, error)

	ValidateSystemTests(params *ValidateSystemTestsParams, opts ...ClientOption) (*ValidateSystemTestsOK, error)

	Func1(params *Func1Params, opts ...ClientOption) (*Func1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CommitFilesToSourceControlSystem commits files to source control associated with a system
*/
func (a *Client) CommitFilesToSourceControlSystem(params *CommitFilesToSourceControlSystemParams, opts ...ClientOption) (*CommitFilesToSourceControlSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitFilesToSourceControlSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CommitFilesToSourceControlSystem",
		Method:             "POST",
		PathPattern:        "/v1/systems/{id}/commits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommitFilesToSourceControlSystemReader{formats: a.formats},
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
	success, ok := result.(*CommitFilesToSourceControlSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CommitFilesToSourceControlSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateSystem creates a system
*/
func (a *Client) CreateSystem(params *CreateSystemParams, opts ...ClientOption) (*CreateSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSystem",
		Method:             "POST",
		PathPattern:        "/v1/systems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSystemReader{formats: a.formats},
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
	success, ok := result.(*CreateSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSystem deletes a system
*/
func (a *Client) DeleteSystem(params *DeleteSystemParams, opts ...ClientOption) (*DeleteSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSystem",
		Method:             "DELETE",
		PathPattern:        "/v1/systems/{system}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSystemReader{formats: a.formats},
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
	success, ok := result.(*DeleteSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUserBranchSystem deletes a user owned branch
*/
func (a *Client) DeleteUserBranchSystem(params *DeleteUserBranchSystemParams, opts ...ClientOption) (*DeleteUserBranchSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserBranchSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserBranchSystem",
		Method:             "DELETE",
		PathPattern:        "/v1/systems/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserBranchSystemReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserBranchSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUserBranchSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAsset gets system asset
*/
func (a *Client) GetAsset(params *GetAssetParams, writer io.Writer, opts ...ClientOption) (*GetAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAsset",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/assets/{assettype}",
		ProducesMediaTypes: []string{"application/gzip", "application/json", "application/x-yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAssetReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*GetAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDefaultPolicies gets default system policies
*/
func (a *Client) GetDefaultPolicies(params *GetDefaultPoliciesParams, opts ...ClientOption) (*GetDefaultPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDefaultPolicies",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/default-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefaultPoliciesReader{formats: a.formats},
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
	success, ok := result.(*GetDefaultPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDefaultPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDefaultPolicy gets default system policy
*/
func (a *Client) GetDefaultPolicy(params *GetDefaultPolicyParams, opts ...ClientOption) (*GetDefaultPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDefaultPolicy",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/default-policies/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefaultPolicyReader{formats: a.formats},
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
	success, ok := result.(*GetDefaultPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDefaultPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInstructions gets system install uninstall instructions
*/
func (a *Client) GetInstructions(params *GetInstructionsParams, opts ...ClientOption) (*GetInstructionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstructionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInstructions",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/instructions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInstructionsReader{formats: a.formats},
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
	success, ok := result.(*GetInstructionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInstructions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOPADiscoveryConfig gets the o p a discovery config for a system
*/
func (a *Client) GetOPADiscoveryConfig(params *GetOPADiscoveryConfigParams, opts ...ClientOption) (*GetOPADiscoveryConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOPADiscoveryConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOPADiscoveryConfig",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/discovery",
		ProducesMediaTypes: []string{"application/gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOPADiscoveryConfigReader{formats: a.formats},
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
	success, ok := result.(*GetOPADiscoveryConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOPADiscoveryConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesBranchSystem gets the list of files for the styra d a s created branch

  Gets the list of files for the branch that the Styra DAS creates when modifying rego in the Styra DAS UI and pushing the changes to GitHub in a branch for review.
*/
func (a *Client) GetSourceControlFilesBranchSystem(params *GetSourceControlFilesBranchSystemParams, opts ...ClientOption) (*GetSourceControlFilesBranchSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesBranchSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesBranchSystem",
		Method:             "GET",
		PathPattern:        "/v1/systems/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSourceControlFilesBranchSystemReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesBranchSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesBranchSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesMasterSystem gets the list of files in the currently chosen branch
*/
func (a *Client) GetSourceControlFilesMasterSystem(params *GetSourceControlFilesMasterSystemParams, opts ...ClientOption) (*GetSourceControlFilesMasterSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesMasterSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesMasterSystem",
		Method:             "GET",
		PathPattern:        "/v1/systems/{id}/master",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSourceControlFilesMasterSystemReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesMasterSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesMasterSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSystem gets a system
*/
func (a *Client) GetSystem(params *GetSystemParams, opts ...ClientOption) (*GetSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSystem",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemReader{formats: a.formats},
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
	success, ok := result.(*GetSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSystemAgents gets system agents
*/
func (a *Client) GetSystemAgents(params *GetSystemAgentsParams, opts ...ClientOption) (*GetSystemAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemAgentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSystemAgents",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemAgentsReader{formats: a.formats},
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
	success, ok := result.(*GetSystemAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSystemAgents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSystemBundle gets system bundle
*/
func (a *Client) GetSystemBundle(params *GetSystemBundleParams, opts ...ClientOption) (*GetSystemBundleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemBundleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSystemBundle",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/bundles/{bundle}/{version}/bundle",
		ProducesMediaTypes: []string{"application/gzip", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemBundleReader{formats: a.formats},
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
	success, ok := result.(*GetSystemBundleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSystemBundle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSystemBundleDeploy gets a system bundle deployment and build status
*/
func (a *Client) GetSystemBundleDeploy(params *GetSystemBundleDeployParams, opts ...ClientOption) (*GetSystemBundleDeployOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemBundleDeployParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSystemBundleDeploy",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/bundle-deploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemBundleDeployReader{formats: a.formats},
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
	success, ok := result.(*GetSystemBundleDeployOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSystemBundleDeploy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSystemBundleDetails gets system bundle details
*/
func (a *Client) GetSystemBundleDetails(params *GetSystemBundleDetailsParams, opts ...ClientOption) (*GetSystemBundleDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemBundleDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSystemBundleDetails",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/bundles/{bundle}/{version}/details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemBundleDetailsReader{formats: a.formats},
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
	success, ok := result.(*GetSystemBundleDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSystemBundleDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSystemBundles lists system bundles starting from the newest towards the oldest
*/
func (a *Client) GetSystemBundles(params *GetSystemBundlesParams, opts ...ClientOption) (*GetSystemBundlesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemBundlesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSystemBundles",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/bundles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemBundlesReader{formats: a.formats},
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
	success, ok := result.(*GetSystemBundlesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSystemBundles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HandleSystemMetrics handles system metrics
*/
func (a *Client) HandleSystemMetrics(params *HandleSystemMetricsParams, opts ...ClientOption) (*HandleSystemMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHandleSystemMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HandleSystemMetrics",
		Method:             "GET",
		PathPattern:        "/v1/systems/metrics",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HandleSystemMetricsReader{formats: a.formats},
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
	success, ok := result.(*HandleSystemMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HandleSystemMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSystems lists systems
*/
func (a *Client) ListSystems(params *ListSystemsParams, opts ...ClientOption) (*ListSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSystemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListSystems",
		Method:             "GET",
		PathPattern:        "/v1/systems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSystemsReader{formats: a.formats},
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
	success, ok := result.(*ListSystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListSystems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RuleSuggestions gets rule suggestions
*/
func (a *Client) RuleSuggestions(params *RuleSuggestionsParams, opts ...ClientOption) (*RuleSuggestionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRuleSuggestionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RuleSuggestions",
		Method:             "GET",
		PathPattern:        "/v1/systems/{system}/suggestions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RuleSuggestionsReader{formats: a.formats},
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
	success, ok := result.(*RuleSuggestionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RuleSuggestions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TranslateExternalIds translates external identifiers to system identifiers
*/
func (a *Client) TranslateExternalIds(params *TranslateExternalIdsParams, opts ...ClientOption) (*TranslateExternalIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTranslateExternalIdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TranslateExternalIds",
		Method:             "POST",
		PathPattern:        "/v1/systems/external-ids",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TranslateExternalIdsReader{formats: a.formats},
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
	success, ok := result.(*TranslateExternalIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TranslateExternalIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSystem creates or update a system
*/
func (a *Client) UpdateSystem(params *UpdateSystemParams, opts ...ClientOption) (*UpdateSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSystem",
		Method:             "PUT",
		PathPattern:        "/v1/systems/{system}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSystemReader{formats: a.formats},
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
	success, ok := result.(*UpdateSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSystemBundleDeploy deploys a system bundle
*/
func (a *Client) UpdateSystemBundleDeploy(params *UpdateSystemBundleDeployParams, opts ...ClientOption) (*UpdateSystemBundleDeployOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSystemBundleDeployParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSystemBundleDeploy",
		Method:             "PUT",
		PathPattern:        "/v1/systems/{system}/bundle-deploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSystemBundleDeployReader{formats: a.formats},
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
	success, ok := result.(*UpdateSystemBundleDeployOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateSystemBundleDeploy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateSystemCompliance validates system compliance
*/
func (a *Client) ValidateSystemCompliance(params *ValidateSystemComplianceParams, opts ...ClientOption) (*ValidateSystemComplianceOK, *ValidateSystemComplianceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateSystemComplianceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ValidateSystemCompliance",
		Method:             "POST",
		PathPattern:        "/v1/systems/{system}/validate/compliance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateSystemComplianceReader{formats: a.formats},
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
	case *ValidateSystemComplianceOK:
		return value, nil, nil
	case *ValidateSystemComplianceAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for systems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateSystemTests validates system unit tests
*/
func (a *Client) ValidateSystemTests(params *ValidateSystemTestsParams, opts ...ClientOption) (*ValidateSystemTestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateSystemTestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ValidateSystemTests",
		Method:             "POST",
		PathPattern:        "/v1/systems/{system}/validate/tests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateSystemTestsReader{formats: a.formats},
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
	success, ok := result.(*ValidateSystemTestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ValidateSystemTests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Func1 verifies git configuration

  verifies that the repository can be accessed with the provided credentials
*/
func (a *Client) Func1(params *Func1Params, opts ...ClientOption) (*Func1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFunc1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "func1",
		Method:             "POST",
		PathPattern:        "/v1/systems/source-control/verify-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Func1Reader{formats: a.formats},
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
	success, ok := result.(*Func1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for func1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
