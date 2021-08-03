// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authz API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authz API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CheckPermissions(params *CheckPermissionsParams, opts ...ClientOption) (*CheckPermissionsOK, error)

	DeleteRoleBinding(params *DeleteRoleBindingParams, opts ...ClientOption) (*DeleteRoleBindingOK, error)

	GetMigrationMarker(params *GetMigrationMarkerParams, opts ...ClientOption) (*GetMigrationMarkerOK, error)

	GetRoleBinding(params *GetRoleBindingParams, opts ...ClientOption) (*GetRoleBindingOK, error)

	ListAllRoleBindings(params *ListAllRoleBindingsParams, opts ...ClientOption) (*ListAllRoleBindingsOK, error)

	ListRoleBindings(params *ListRoleBindingsParams, opts ...ClientOption) (*ListRoleBindingsOK, error)

	ListRoles(params *ListRolesParams, opts ...ClientOption) (*ListRolesOK, error)

	QueryRoleBindings(params *QueryRoleBindingsParams, opts ...ClientOption) (*QueryRoleBindingsOK, error)

	UpdateMigrationMarker(params *UpdateMigrationMarkerParams, opts ...ClientOption) (*UpdateMigrationMarkerOK, error)

	UpdateRoleBinding(params *UpdateRoleBindingParams, opts ...ClientOption) (*UpdateRoleBindingOK, error)

	UpdateRoleBindings(params *UpdateRoleBindingsParams, opts ...ClientOption) (*UpdateRoleBindingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CheckPermissions evaluates a list of permissions
*/
func (a *Client) CheckPermissions(params *CheckPermissionsParams, opts ...ClientOption) (*CheckPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CheckPermissions",
		Method:             "POST",
		PathPattern:        "/v1/authz/evaluation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CheckPermissionsReader{formats: a.formats},
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
	success, ok := result.(*CheckPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CheckPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRoleBinding deletes a resource role biding
*/
func (a *Client) DeleteRoleBinding(params *DeleteRoleBindingParams, opts ...ClientOption) (*DeleteRoleBindingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleBindingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRoleBinding",
		Method:             "DELETE",
		PathPattern:        "/v1/authz/rolebindings/{resourcetype}/{resource}/{rolebinding}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoleBindingReader{formats: a.formats},
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
	success, ok := result.(*DeleteRoleBindingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteRoleBinding: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMigrationMarker gets migration marker
*/
func (a *Client) GetMigrationMarker(params *GetMigrationMarkerParams, opts ...ClientOption) (*GetMigrationMarkerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMigrationMarkerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMigrationMarker",
		Method:             "GET",
		PathPattern:        "/v2/authz/rolebindings/marker",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMigrationMarkerReader{formats: a.formats},
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
	success, ok := result.(*GetMigrationMarkerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMigrationMarker: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRoleBinding gets a role binding
*/
func (a *Client) GetRoleBinding(params *GetRoleBindingParams, opts ...ClientOption) (*GetRoleBindingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleBindingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRoleBinding",
		Method:             "GET",
		PathPattern:        "/v1/authz/rolebindings/{resourcetype}/{resource}/{rolebinding}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRoleBindingReader{formats: a.formats},
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
	success, ok := result.(*GetRoleBindingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRoleBinding: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllRoleBindings lists all role bindings for all resources of all resource types
*/
func (a *Client) ListAllRoleBindings(params *ListAllRoleBindingsParams, opts ...ClientOption) (*ListAllRoleBindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllRoleBindingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListAllRoleBindings",
		Method:             "GET",
		PathPattern:        "/v1/authz/rolebindings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAllRoleBindingsReader{formats: a.formats},
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
	success, ok := result.(*ListAllRoleBindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListAllRoleBindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRoleBindings lists role bindings
*/
func (a *Client) ListRoleBindings(params *ListRoleBindingsParams, opts ...ClientOption) (*ListRoleBindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRoleBindingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListRoleBindings",
		Method:             "GET",
		PathPattern:        "/v2/authz/rolebindings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRoleBindingsReader{formats: a.formats},
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
	success, ok := result.(*ListRoleBindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListRoleBindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRoles lists styra defined roles
*/
func (a *Client) ListRoles(params *ListRolesParams, opts ...ClientOption) (*ListRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListRoles",
		Method:             "GET",
		PathPattern:        "/v2/authz/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRolesReader{formats: a.formats},
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
	success, ok := result.(*ListRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryRoleBindings queries role bindings
*/
func (a *Client) QueryRoleBindings(params *QueryRoleBindingsParams, opts ...ClientOption) (*QueryRoleBindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryRoleBindingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryRoleBindings",
		Method:             "POST",
		PathPattern:        "/v2/authz/rolebindings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryRoleBindingsReader{formats: a.formats},
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
	success, ok := result.(*QueryRoleBindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryRoleBindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateMigrationMarker updates migration marker
*/
func (a *Client) UpdateMigrationMarker(params *UpdateMigrationMarkerParams, opts ...ClientOption) (*UpdateMigrationMarkerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMigrationMarkerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateMigrationMarker",
		Method:             "PUT",
		PathPattern:        "/v2/authz/rolebindings/marker",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMigrationMarkerReader{formats: a.formats},
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
	success, ok := result.(*UpdateMigrationMarkerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateMigrationMarker: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRoleBinding updates a role binding
*/
func (a *Client) UpdateRoleBinding(params *UpdateRoleBindingParams, opts ...ClientOption) (*UpdateRoleBindingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleBindingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRoleBinding",
		Method:             "PUT",
		PathPattern:        "/v1/authz/rolebindings/{resourcetype}/{resource}/{rolebinding}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRoleBindingReader{formats: a.formats},
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
	success, ok := result.(*UpdateRoleBindingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateRoleBinding: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRoleBindings updates role bindings
*/
func (a *Client) UpdateRoleBindings(params *UpdateRoleBindingsParams, opts ...ClientOption) (*UpdateRoleBindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleBindingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRoleBindings",
		Method:             "PUT",
		PathPattern:        "/v2/authz/rolebindings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRoleBindingsReader{formats: a.formats},
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
	success, ok := result.(*UpdateRoleBindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateRoleBindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
