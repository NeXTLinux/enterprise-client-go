/*
Nextlinux Enterprise API Server

This is the Nextlinux Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
Contact: dev@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external

import (
	"encoding/json"
)

// InventoryClusterConfig Cluster specific configuration
type InventoryClusterConfig struct {
	CredentialType *string `json:"credential_type,omitempty"`
	Namespaces []string `json:"namespaces,omitempty"`
	// FQDN of the cluster API server
	ClusterServer *string `json:"cluster_server,omitempty"`
	// Base64 Encoded Public Certificate for the cluster
	ClusterCertificate *string `json:"cluster_certificate,omitempty"`
	// Base64 Encoded Public Certificate for the client. Not required if credential_type == token
	ClientCertificate *string `json:"client_certificate,omitempty"`
	// Base64 Encoded credential for the client
	Credential *string `json:"credential,omitempty"`
}

// NewInventoryClusterConfig instantiates a new InventoryClusterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryClusterConfig() *InventoryClusterConfig {
	this := InventoryClusterConfig{}
	return &this
}

// NewInventoryClusterConfigWithDefaults instantiates a new InventoryClusterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryClusterConfigWithDefaults() *InventoryClusterConfig {
	this := InventoryClusterConfig{}
	return &this
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise.
func (o *InventoryClusterConfig) GetCredentialType() string {
	if o == nil || o.CredentialType == nil {
		var ret string
		return ret
	}
	return *o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryClusterConfig) GetCredentialTypeOk() (*string, bool) {
	if o == nil || o.CredentialType == nil {
		return nil, false
	}
	return o.CredentialType, true
}

// HasCredentialType returns a boolean if a field has been set.
func (o *InventoryClusterConfig) HasCredentialType() bool {
	if o != nil && o.CredentialType != nil {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given string and assigns it to the CredentialType field.
func (o *InventoryClusterConfig) SetCredentialType(v string) {
	o.CredentialType = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *InventoryClusterConfig) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryClusterConfig) GetNamespacesOk() ([]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *InventoryClusterConfig) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *InventoryClusterConfig) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetClusterServer returns the ClusterServer field value if set, zero value otherwise.
func (o *InventoryClusterConfig) GetClusterServer() string {
	if o == nil || o.ClusterServer == nil {
		var ret string
		return ret
	}
	return *o.ClusterServer
}

// GetClusterServerOk returns a tuple with the ClusterServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryClusterConfig) GetClusterServerOk() (*string, bool) {
	if o == nil || o.ClusterServer == nil {
		return nil, false
	}
	return o.ClusterServer, true
}

// HasClusterServer returns a boolean if a field has been set.
func (o *InventoryClusterConfig) HasClusterServer() bool {
	if o != nil && o.ClusterServer != nil {
		return true
	}

	return false
}

// SetClusterServer gets a reference to the given string and assigns it to the ClusterServer field.
func (o *InventoryClusterConfig) SetClusterServer(v string) {
	o.ClusterServer = &v
}

// GetClusterCertificate returns the ClusterCertificate field value if set, zero value otherwise.
func (o *InventoryClusterConfig) GetClusterCertificate() string {
	if o == nil || o.ClusterCertificate == nil {
		var ret string
		return ret
	}
	return *o.ClusterCertificate
}

// GetClusterCertificateOk returns a tuple with the ClusterCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryClusterConfig) GetClusterCertificateOk() (*string, bool) {
	if o == nil || o.ClusterCertificate == nil {
		return nil, false
	}
	return o.ClusterCertificate, true
}

// HasClusterCertificate returns a boolean if a field has been set.
func (o *InventoryClusterConfig) HasClusterCertificate() bool {
	if o != nil && o.ClusterCertificate != nil {
		return true
	}

	return false
}

// SetClusterCertificate gets a reference to the given string and assigns it to the ClusterCertificate field.
func (o *InventoryClusterConfig) SetClusterCertificate(v string) {
	o.ClusterCertificate = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *InventoryClusterConfig) GetClientCertificate() string {
	if o == nil || o.ClientCertificate == nil {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryClusterConfig) GetClientCertificateOk() (*string, bool) {
	if o == nil || o.ClientCertificate == nil {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *InventoryClusterConfig) HasClientCertificate() bool {
	if o != nil && o.ClientCertificate != nil {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *InventoryClusterConfig) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *InventoryClusterConfig) GetCredential() string {
	if o == nil || o.Credential == nil {
		var ret string
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryClusterConfig) GetCredentialOk() (*string, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *InventoryClusterConfig) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given string and assigns it to the Credential field.
func (o *InventoryClusterConfig) SetCredential(v string) {
	o.Credential = &v
}

func (o InventoryClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialType != nil {
		toSerialize["credential_type"] = o.CredentialType
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.ClusterServer != nil {
		toSerialize["cluster_server"] = o.ClusterServer
	}
	if o.ClusterCertificate != nil {
		toSerialize["cluster_certificate"] = o.ClusterCertificate
	}
	if o.ClientCertificate != nil {
		toSerialize["client_certificate"] = o.ClientCertificate
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryClusterConfig struct {
	value *InventoryClusterConfig
	isSet bool
}

func (v NullableInventoryClusterConfig) Get() *InventoryClusterConfig {
	return v.value
}

func (v *NullableInventoryClusterConfig) Set(val *InventoryClusterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryClusterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryClusterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryClusterConfig(val *InventoryClusterConfig) *NullableInventoryClusterConfig {
	return &NullableInventoryClusterConfig{value: val, isSet: true}
}

func (v NullableInventoryClusterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryClusterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


