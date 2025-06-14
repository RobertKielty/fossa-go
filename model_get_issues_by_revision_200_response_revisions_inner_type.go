/*
FOSSA API

OpenAPI Specification for public FOSSA APIs

API version: 4.28.61
Contact: support@fossa.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fossa

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// GetIssuesByRevision200ResponseRevisionsInnerType - struct for GetIssuesByRevision200ResponseRevisionsInnerType
type GetIssuesByRevision200ResponseRevisionsInnerType struct {
	GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf
	GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1
	GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2
}

// GetIssuesByRevision200ResponseRevisionsInnerTypeOneOfAsGetIssuesByRevision200ResponseRevisionsInnerType is a convenience function that returns GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf wrapped in GetIssuesByRevision200ResponseRevisionsInnerType
func GetIssuesByRevision200ResponseRevisionsInnerTypeOneOfAsGetIssuesByRevision200ResponseRevisionsInnerType(v *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf) GetIssuesByRevision200ResponseRevisionsInnerType {
	return GetIssuesByRevision200ResponseRevisionsInnerType{
		GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf: v,
	}
}

// GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1AsGetIssuesByRevision200ResponseRevisionsInnerType is a convenience function that returns GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 wrapped in GetIssuesByRevision200ResponseRevisionsInnerType
func GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1AsGetIssuesByRevision200ResponseRevisionsInnerType(v *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1) GetIssuesByRevision200ResponseRevisionsInnerType {
	return GetIssuesByRevision200ResponseRevisionsInnerType{
		GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1: v,
	}
}

// GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2AsGetIssuesByRevision200ResponseRevisionsInnerType is a convenience function that returns GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 wrapped in GetIssuesByRevision200ResponseRevisionsInnerType
func GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2AsGetIssuesByRevision200ResponseRevisionsInnerType(v *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetIssuesByRevision200ResponseRevisionsInnerType {
	return GetIssuesByRevision200ResponseRevisionsInnerType{
		GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetIssuesByRevision200ResponseRevisionsInnerType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf
	err = newStrictDecoder(data).Decode(&dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf)
	if err == nil {
		jsonGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf, _ := json.Marshal(dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf)
		if string(jsonGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf) == "{}" { // empty struct
			dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf = nil
		} else {
			if err = validator.Validate(dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf); err != nil {
				dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf = nil
	}

	// try to unmarshal data into GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1
	err = newStrictDecoder(data).Decode(&dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1)
	if err == nil {
		jsonGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1, _ := json.Marshal(dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1)
		if string(jsonGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1) == "{}" { // empty struct
			dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 = nil
		} else {
			if err = validator.Validate(dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1); err != nil {
				dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 = nil
	}

	// try to unmarshal data into GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2
	err = newStrictDecoder(data).Decode(&dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2)
	if err == nil {
		jsonGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2, _ := json.Marshal(dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2)
		if string(jsonGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) == "{}" { // empty struct
			dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 = nil
		} else {
			if err = validator.Validate(dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2); err != nil {
				dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf = nil
		dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 = nil
		dst.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetIssuesByRevision200ResponseRevisionsInnerType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetIssuesByRevision200ResponseRevisionsInnerType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetIssuesByRevision200ResponseRevisionsInnerType) MarshalJSON() ([]byte, error) {
	if src.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf != nil {
		return json.Marshal(&src.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf)
	}

	if src.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 != nil {
		return json.Marshal(&src.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1)
	}

	if src.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 != nil {
		return json.Marshal(&src.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetIssuesByRevision200ResponseRevisionsInnerType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf != nil {
		return obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf
	}

	if obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 != nil {
		return obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1
	}

	if obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 != nil {
		return obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetIssuesByRevision200ResponseRevisionsInnerType) GetActualInstanceValue() (interface{}) {
	if obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf != nil {
		return *obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf
	}

	if obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1 != nil {
		return *obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf1
	}

	if obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 != nil {
		return *obj.GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableGetIssuesByRevision200ResponseRevisionsInnerType struct {
	value *GetIssuesByRevision200ResponseRevisionsInnerType
	isSet bool
}

func (v NullableGetIssuesByRevision200ResponseRevisionsInnerType) Get() *GetIssuesByRevision200ResponseRevisionsInnerType {
	return v.value
}

func (v *NullableGetIssuesByRevision200ResponseRevisionsInnerType) Set(val *GetIssuesByRevision200ResponseRevisionsInnerType) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssuesByRevision200ResponseRevisionsInnerType) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssuesByRevision200ResponseRevisionsInnerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssuesByRevision200ResponseRevisionsInnerType(val *GetIssuesByRevision200ResponseRevisionsInnerType) *NullableGetIssuesByRevision200ResponseRevisionsInnerType {
	return &NullableGetIssuesByRevision200ResponseRevisionsInnerType{value: val, isSet: true}
}

func (v NullableGetIssuesByRevision200ResponseRevisionsInnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssuesByRevision200ResponseRevisionsInnerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


