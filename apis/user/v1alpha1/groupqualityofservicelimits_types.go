/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"errors"
	"reflect"

	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/utils/ptr"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// +kubebuilder:validation:Pattern=`^(0|((0|[1-9][0-9]*)[KMGT]i))$`
type Quantity string

func (q *Quantity) ToKiB() (*int64, error) {
	if q == nil {
		return ptr.To(int64(-1)), nil
	}

	rq, err := resource.ParseQuantity(string(*q))
	if err != nil {
		return nil, err
	}

	i, ok := rq.AsInt64()
	if !ok {
		i, ok = rq.AsDec().Unscaled()
		if !ok {
			return nil, errors.New("Unable to convert Quantity to KiB int")
		}
	}

	i /= 1024
	return &i, nil
}

// QualityOfService configures data limits. The value -1 indicates unlimited.
type QualityOfServiceLimits struct {
	// StorageQuotaBytes is the limit for total stored data in KiB.
	//+optional
	StorageQuotaBytes *Quantity `json:"storageQuotaBytes,omitempty"`
	// StorageQuotaCount is the limit for total number of objects.
	//+optional
	StorageQuotaCount *uint32 `json:"storageQuotaCount,omitempty"`
	// RequestsPerMin is the limit for number of HTTP requests per minute.
	//+optional
	RequestsPerMin *uint32 `json:"requestsPerMin,omitempty"`
	// InboundBytesPerMin is the limit for inbound data per minute in KiB.
	//+optional
	InboundBytesPerMin *Quantity `json:"inboundBytesPerMin,omitempty"`
	// OutboundKiBsPerMin is the limit for outbound data per minute in KiB.
	//+optional
	OutboundBytesPerMin *Quantity `json:"outboundBytesPerMin,omitempty"`
}

// GroupQualityOfServiceLimitsParameters are the configurable fields of a GroupQualityOfServiceLimits.
type GroupQualityOfServiceLimitsParameters struct {
	// Group for the quality of service limits.
	//+optional
	//+immutable
	GroupID string `json:"groupId,omitempty"`

	// GroupIDRef is a reference to a group to retrieve its groupId.
	//+optional
	//+immutable
	GroupIDRef *xpv1.Reference `json:"groupIdRef,omitempty"`

	// GroupIDSelector selects reference to a group to retrieve its groupId.
	//+optional
	GroupIDSelector *xpv1.Selector `json:"groupIdSelector,omitempty"`

	// Region in which to apply the quality of service limits. Default region if unspecified.
	Region string `json:"region,omitempty"`

	// Warning is the soft limit that triggers a warning.
	//+optional
	Warning *QualityOfServiceLimits `json:"warning,omitempty"`

	// Hard is the hard limit.
	//+optional
	Hard *QualityOfServiceLimits `json:"hard,omitempty"`
}

// GroupQualityOfServiceLimitsObservation are the observable fields of a GroupQualityOfServiceLimits.
type GroupQualityOfServiceLimitsObservation struct {
}

// A GroupQualityOfServiceLimitsSpec defines the desired state of a GroupQualityOfServiceLimits.
type GroupQualityOfServiceLimitsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GroupQualityOfServiceLimitsParameters `json:"forProvider"`
}

// A GroupQualityOfServiceLimitsStatus represents the observed state of a GroupQualityOfServiceLimits.
type GroupQualityOfServiceLimitsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GroupQualityOfServiceLimitsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A GroupQualityOfServiceLimits is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudian}
type GroupQualityOfServiceLimits struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GroupQualityOfServiceLimitsSpec   `json:"spec"`
	Status GroupQualityOfServiceLimitsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupQualityOfServiceLimitsList contains a list of GroupQualityOfServiceLimits
type GroupQualityOfServiceLimitsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupQualityOfServiceLimits `json:"items"`
}

// GroupQualityOfServiceLimits type metadata.
var (
	GroupQualityOfServiceLimitsKind             = reflect.TypeOf(GroupQualityOfServiceLimits{}).Name()
	GroupQualityOfServiceLimitsGroupKind        = schema.GroupKind{Group: MetadataGroup, Kind: GroupQualityOfServiceLimitsKind}.String()
	GroupQualityOfServiceLimitsKindAPIVersion   = GroupQualityOfServiceLimitsKind + "." + SchemeGroupVersion.String()
	GroupQualityOfServiceLimitsGroupVersionKind = SchemeGroupVersion.WithKind(GroupQualityOfServiceLimitsKind)
)

func init() {
	SchemeBuilder.Register(&GroupQualityOfServiceLimits{}, &GroupQualityOfServiceLimitsList{})
}
