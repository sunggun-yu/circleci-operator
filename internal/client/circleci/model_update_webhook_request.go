/*
CircleCI API

This describes the resources that make up the CircleCI API v2.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package circleci

import (
	"encoding/json"
)

// checks if the UpdateWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebhookRequest{}

// UpdateWebhookRequest The parameters for an update webhook request
type UpdateWebhookRequest struct {
	// Name of the webhook
	Name *string `json:"name,omitempty"`
	// Events that will trigger the webhook
	Events []string `json:"events,omitempty"`
	// URL to deliver the webhook to. Note: protocol must be included as well (only https is supported)
	Url *string `json:"url,omitempty"`
	// Secret used to build an HMAC hash of the payload and passed as a header in the webhook request
	SigningSecret *string `json:"signing-secret,omitempty"`
	// Whether to enforce TLS certificate verification when delivering the webhook
	VerifyTls *bool `json:"verify-tls,omitempty"`
}

// NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhookRequest() *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	return &this
}

// NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateWebhookRequest) SetName(v string) {
	o.Name = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetEvents() []string {
	if o == nil || IsNil(o.Events) {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *UpdateWebhookRequest) SetEvents(v []string) {
	o.Events = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateWebhookRequest) SetUrl(v string) {
	o.Url = &v
}

// GetSigningSecret returns the SigningSecret field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetSigningSecret() string {
	if o == nil || IsNil(o.SigningSecret) {
		var ret string
		return ret
	}
	return *o.SigningSecret
}

// GetSigningSecretOk returns a tuple with the SigningSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetSigningSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SigningSecret) {
		return nil, false
	}
	return o.SigningSecret, true
}

// HasSigningSecret returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasSigningSecret() bool {
	if o != nil && !IsNil(o.SigningSecret) {
		return true
	}

	return false
}

// SetSigningSecret gets a reference to the given string and assigns it to the SigningSecret field.
func (o *UpdateWebhookRequest) SetSigningSecret(v string) {
	o.SigningSecret = &v
}

// GetVerifyTls returns the VerifyTls field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetVerifyTls() bool {
	if o == nil || IsNil(o.VerifyTls) {
		var ret bool
		return ret
	}
	return *o.VerifyTls
}

// GetVerifyTlsOk returns a tuple with the VerifyTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetVerifyTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyTls) {
		return nil, false
	}
	return o.VerifyTls, true
}

// HasVerifyTls returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasVerifyTls() bool {
	if o != nil && !IsNil(o.VerifyTls) {
		return true
	}

	return false
}

// SetVerifyTls gets a reference to the given bool and assigns it to the VerifyTls field.
func (o *UpdateWebhookRequest) SetVerifyTls(v bool) {
	o.VerifyTls = &v
}

func (o UpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.SigningSecret) {
		toSerialize["signing-secret"] = o.SigningSecret
	}
	if !IsNil(o.VerifyTls) {
		toSerialize["verify-tls"] = o.VerifyTls
	}
	return toSerialize, nil
}

type NullableUpdateWebhookRequest struct {
	value *UpdateWebhookRequest
	isSet bool
}

func (v NullableUpdateWebhookRequest) Get() *UpdateWebhookRequest {
	return v.value
}

func (v *NullableUpdateWebhookRequest) Set(val *UpdateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhookRequest(val *UpdateWebhookRequest) *NullableUpdateWebhookRequest {
	return &NullableUpdateWebhookRequest{value: val, isSet: true}
}

func (v NullableUpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
