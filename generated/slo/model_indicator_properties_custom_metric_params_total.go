/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the IndicatorPropertiesCustomMetricParamsTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicatorPropertiesCustomMetricParamsTotal{}

// IndicatorPropertiesCustomMetricParamsTotal An object defining the \"total\" metrics and equation
type IndicatorPropertiesCustomMetricParamsTotal struct {
	// List of metrics with their name, aggregation type, and field.
	Metrics []IndicatorPropertiesCustomMetricParamsTotalMetricsInner `json:"metrics"`
	// The equation to calculate the \"total\" metric.
	Equation string `json:"equation"`
}

// NewIndicatorPropertiesCustomMetricParamsTotal instantiates a new IndicatorPropertiesCustomMetricParamsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorPropertiesCustomMetricParamsTotal(metrics []IndicatorPropertiesCustomMetricParamsTotalMetricsInner, equation string) *IndicatorPropertiesCustomMetricParamsTotal {
	this := IndicatorPropertiesCustomMetricParamsTotal{}
	this.Metrics = metrics
	this.Equation = equation
	return &this
}

// NewIndicatorPropertiesCustomMetricParamsTotalWithDefaults instantiates a new IndicatorPropertiesCustomMetricParamsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorPropertiesCustomMetricParamsTotalWithDefaults() *IndicatorPropertiesCustomMetricParamsTotal {
	this := IndicatorPropertiesCustomMetricParamsTotal{}
	return &this
}

// GetMetrics returns the Metrics field value
func (o *IndicatorPropertiesCustomMetricParamsTotal) GetMetrics() []IndicatorPropertiesCustomMetricParamsTotalMetricsInner {
	if o == nil {
		var ret []IndicatorPropertiesCustomMetricParamsTotalMetricsInner
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesCustomMetricParamsTotal) GetMetricsOk() ([]IndicatorPropertiesCustomMetricParamsTotalMetricsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *IndicatorPropertiesCustomMetricParamsTotal) SetMetrics(v []IndicatorPropertiesCustomMetricParamsTotalMetricsInner) {
	o.Metrics = v
}

// GetEquation returns the Equation field value
func (o *IndicatorPropertiesCustomMetricParamsTotal) GetEquation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Equation
}

// GetEquationOk returns a tuple with the Equation field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesCustomMetricParamsTotal) GetEquationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Equation, true
}

// SetEquation sets field value
func (o *IndicatorPropertiesCustomMetricParamsTotal) SetEquation(v string) {
	o.Equation = v
}

func (o IndicatorPropertiesCustomMetricParamsTotal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicatorPropertiesCustomMetricParamsTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metrics"] = o.Metrics
	toSerialize["equation"] = o.Equation
	return toSerialize, nil
}

type NullableIndicatorPropertiesCustomMetricParamsTotal struct {
	value *IndicatorPropertiesCustomMetricParamsTotal
	isSet bool
}

func (v NullableIndicatorPropertiesCustomMetricParamsTotal) Get() *IndicatorPropertiesCustomMetricParamsTotal {
	return v.value
}

func (v *NullableIndicatorPropertiesCustomMetricParamsTotal) Set(val *IndicatorPropertiesCustomMetricParamsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorPropertiesCustomMetricParamsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorPropertiesCustomMetricParamsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorPropertiesCustomMetricParamsTotal(val *IndicatorPropertiesCustomMetricParamsTotal) *NullableIndicatorPropertiesCustomMetricParamsTotal {
	return &NullableIndicatorPropertiesCustomMetricParamsTotal{value: val, isSet: true}
}

func (v NullableIndicatorPropertiesCustomMetricParamsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorPropertiesCustomMetricParamsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
