/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// CustomResourceColumnDefinitionApplyConfiguration represents an declarative configuration of the CustomResourceColumnDefinition type for use
// with apply.
type CustomResourceColumnDefinitionApplyConfiguration struct {
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Format      *string `json:"format,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
	JSONPath    *string `json:"JSONPath,omitempty"`
}

// CustomResourceColumnDefinitionApplyConfiguration constructs an declarative configuration of the CustomResourceColumnDefinition type for use with
// apply.
func CustomResourceColumnDefinition() *CustomResourceColumnDefinitionApplyConfiguration {
	return &CustomResourceColumnDefinitionApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *CustomResourceColumnDefinitionApplyConfiguration) WithName(value string) *CustomResourceColumnDefinitionApplyConfiguration {
	b.Name = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *CustomResourceColumnDefinitionApplyConfiguration) WithType(value string) *CustomResourceColumnDefinitionApplyConfiguration {
	b.Type = &value
	return b
}

// WithFormat sets the Format field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Format field is set to the value of the last call.
func (b *CustomResourceColumnDefinitionApplyConfiguration) WithFormat(value string) *CustomResourceColumnDefinitionApplyConfiguration {
	b.Format = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *CustomResourceColumnDefinitionApplyConfiguration) WithDescription(value string) *CustomResourceColumnDefinitionApplyConfiguration {
	b.Description = &value
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *CustomResourceColumnDefinitionApplyConfiguration) WithPriority(value int32) *CustomResourceColumnDefinitionApplyConfiguration {
	b.Priority = &value
	return b
}

// WithJSONPath sets the JSONPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JSONPath field is set to the value of the last call.
func (b *CustomResourceColumnDefinitionApplyConfiguration) WithJSONPath(value string) *CustomResourceColumnDefinitionApplyConfiguration {
	b.JSONPath = &value
	return b
}
