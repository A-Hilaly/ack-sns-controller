// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TopicSpec defines the desired state of Topic.
//
// A wrapper type for the topic's Amazon Resource Name (ARN). To retrieve a
// topic's attributes, use GetTopicAttributes.
type TopicSpec struct {
	ApplicationFailureFeedbackRoleARN    *string `json:"applicationFailureFeedbackRoleARN,omitempty"`
	ApplicationSuccessFeedbackRoleARN    *string `json:"applicationSuccessFeedbackRoleARN,omitempty"`
	ApplicationSuccessFeedbackSampleRate *string `json:"applicationSuccessFeedbackSampleRate,omitempty"`
	ContentBasedDeduplication            *string `json:"contentBasedDeduplication,omitempty"`
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30,720.
	DataProtectionPolicy              *string                                  `json:"dataProtectionPolicy,omitempty"`
	DeliveryPolicy                    *string                                  `json:"deliveryPolicy,omitempty"`
	DisplayName                       *string                                  `json:"displayName,omitempty"`
	FIFOTopic                         *string                                  `json:"fifoTopic,omitempty"`
	FirehoseFailureFeedbackRoleARN    *string                                  `json:"firehoseFailureFeedbackRoleARN,omitempty"`
	FirehoseSuccessFeedbackRoleARN    *string                                  `json:"firehoseSuccessFeedbackRoleARN,omitempty"`
	FirehoseSuccessFeedbackSampleRate *string                                  `json:"firehoseSuccessFeedbackSampleRate,omitempty"`
	HTTPFailureFeedbackRoleARN        *string                                  `json:"httpFailureFeedbackRoleARN,omitempty"`
	HTTPSuccessFeedbackRoleARN        *string                                  `json:"httpSuccessFeedbackRoleARN,omitempty"`
	HTTPSuccessFeedbackSampleRate     *string                                  `json:"httpSuccessFeedbackSampleRate,omitempty"`
	KMSMasterKeyID                    *string                                  `json:"kmsMasterKeyID,omitempty"`
	KMSMasterKeyRef                   *ackv1alpha1.AWSResourceReferenceWrapper `json:"kmsMasterKeyRef,omitempty"`
	LambdaFailureFeedbackRoleARN      *string                                  `json:"lambdaFailureFeedbackRoleARN,omitempty"`
	LambdaSuccessFeedbackRoleARN      *string                                  `json:"lambdaSuccessFeedbackRoleARN,omitempty"`
	LambdaSuccessFeedbackSampleRate   *string                                  `json:"lambdaSuccessFeedbackSampleRate,omitempty"`
	// The name of the topic you want to create.
	//
	// Constraints: Topic names must be made up of only uppercase and lowercase
	// ASCII letters, numbers, underscores, and hyphens, and must be between 1 and
	// 256 characters long.
	//
	// For a FIFO (first-in-first-out) topic, the name must end with the .fifo suffix.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
	// +kubebuilder:validation:Required
	Name                         *string                                  `json:"name"`
	Policy                       *string                                  `json:"policy,omitempty"`
	PolicyRef                    *ackv1alpha1.AWSResourceReferenceWrapper `json:"policyRef,omitempty"`
	SQSFailureFeedbackRoleARN    *string                                  `json:"sqsFailureFeedbackRoleARN,omitempty"`
	SQSSuccessFeedbackRoleARN    *string                                  `json:"sqsSuccessFeedbackRoleARN,omitempty"`
	SQSSuccessFeedbackSampleRate *string                                  `json:"sqsSuccessFeedbackSampleRate,omitempty"`
	SignatureVersion             *string                                  `json:"signatureVersion,omitempty"`
	// The list of tags to add to a new topic.
	//
	// To be able to tag a topic on creation, you must have the sns:CreateTopic
	// and sns:TagResource permissions.
	Tags          []*Tag  `json:"tags,omitempty"`
	TracingConfig *string `json:"tracingConfig,omitempty"`
}

// TopicStatus defines the observed state of Topic
type TopicStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	EffectiveDeliveryPolicy *string `json:"effectiveDeliveryPolicy,omitempty"`
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty"`
	// +kubebuilder:validation:Optional
	TopicARN *string `json:"topicARN,omitempty"`
}

// Topic is the Schema for the Topics API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ARN",type=string,priority=1,JSONPath=`.status.ackResourceMetadata.arn`
// +kubebuilder:printcolumn:name="DISPLAYNAME",type=string,priority=0,JSONPath=`.spec.displayName`
// +kubebuilder:printcolumn:name="Synced",type="string",priority=0,JSONPath=".status.conditions[?(@.type==\"ACK.ResourceSynced\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",priority=0,JSONPath=".metadata.creationTimestamp"
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec,omitempty"`
	Status            TopicStatus `json:"status,omitempty"`
}

// TopicList contains a list of Topic
// +kubebuilder:object:root=true
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
