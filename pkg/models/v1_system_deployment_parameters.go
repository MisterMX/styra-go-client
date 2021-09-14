// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SystemDeploymentParameters v1 system deployment parameters
//
// swagger:model v1.SystemDeploymentParameters
type V1SystemDeploymentParameters struct {

	// true to fail close
	DenyOnOpaFail *bool `json:"deny_on_opa_fail,omitempty"`

	// true to enable partial eval for the system's bundle'
	EnablePartialEval *bool `json:"enable_partial_eval,omitempty"`

	// extra deployment settings
	Extra interface{} `json:"extra,omitempty"`

	// HTTP proxy URL
	HTTPProxy string `json:"http_proxy,omitempty"`

	// HTTPS proxy URL
	HTTPSProxy string `json:"https_proxy,omitempty"`

	// minimum Kubernetes version expected (where applicable)
	KubernetesVersion string `json:"kubernetes_version,omitempty"`

	// Kubernetes namespace the system is deployed to
	Namespace string `json:"namespace,omitempty"`

	// URLs that should be excluded from proxying
	NoProxy string `json:"no_proxy,omitempty"`

	// Kubernetes webhook timeout (where applicable)
	TimeoutSeconds int32 `json:"timeout_seconds,omitempty"`

	// trusted CA certificates
	TrustedCaCerts []string `json:"trusted_ca_certs"`

	// trusted container registry
	TrustedContainerRegistry string `json:"trusted_container_registry,omitempty"`
}

// Validate validates this v1 system deployment parameters
func (m *V1SystemDeploymentParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 system deployment parameters based on context it is used
func (m *V1SystemDeploymentParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SystemDeploymentParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SystemDeploymentParameters) UnmarshalBinary(b []byte) error {
	var res V1SystemDeploymentParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
