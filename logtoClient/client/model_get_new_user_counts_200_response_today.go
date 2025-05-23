/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetNewUserCounts200ResponseToday type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNewUserCounts200ResponseToday{}

// GetNewUserCounts200ResponseToday struct for GetNewUserCounts200ResponseToday
type GetNewUserCounts200ResponseToday struct {
	Count float32 `json:"count"`
	Delta float32 `json:"delta"`
}

type _GetNewUserCounts200ResponseToday GetNewUserCounts200ResponseToday

// NewGetNewUserCounts200ResponseToday instantiates a new GetNewUserCounts200ResponseToday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNewUserCounts200ResponseToday(count float32, delta float32) *GetNewUserCounts200ResponseToday {
	this := GetNewUserCounts200ResponseToday{}
	this.Count = count
	this.Delta = delta
	return &this
}

// NewGetNewUserCounts200ResponseTodayWithDefaults instantiates a new GetNewUserCounts200ResponseToday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNewUserCounts200ResponseTodayWithDefaults() *GetNewUserCounts200ResponseToday {
	this := GetNewUserCounts200ResponseToday{}
	return &this
}

// GetCount returns the Count field value
func (o *GetNewUserCounts200ResponseToday) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetNewUserCounts200ResponseToday) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetNewUserCounts200ResponseToday) SetCount(v float32) {
	o.Count = v
}

// GetDelta returns the Delta field value
func (o *GetNewUserCounts200ResponseToday) GetDelta() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value
// and a boolean to check if the value has been set.
func (o *GetNewUserCounts200ResponseToday) GetDeltaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delta, true
}

// SetDelta sets field value
func (o *GetNewUserCounts200ResponseToday) SetDelta(v float32) {
	o.Delta = v
}

func (o GetNewUserCounts200ResponseToday) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNewUserCounts200ResponseToday) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["delta"] = o.Delta
	return toSerialize, nil
}

func (o *GetNewUserCounts200ResponseToday) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"delta",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetNewUserCounts200ResponseToday := _GetNewUserCounts200ResponseToday{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetNewUserCounts200ResponseToday)

	if err != nil {
		return err
	}

	*o = GetNewUserCounts200ResponseToday(varGetNewUserCounts200ResponseToday)

	return err
}

type NullableGetNewUserCounts200ResponseToday struct {
	value *GetNewUserCounts200ResponseToday
	isSet bool
}

func (v NullableGetNewUserCounts200ResponseToday) Get() *GetNewUserCounts200ResponseToday {
	return v.value
}

func (v *NullableGetNewUserCounts200ResponseToday) Set(val *GetNewUserCounts200ResponseToday) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNewUserCounts200ResponseToday) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNewUserCounts200ResponseToday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNewUserCounts200ResponseToday(val *GetNewUserCounts200ResponseToday) *NullableGetNewUserCounts200ResponseToday {
	return &NullableGetNewUserCounts200ResponseToday{value: val, isSet: true}
}

func (v NullableGetNewUserCounts200ResponseToday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNewUserCounts200ResponseToday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


