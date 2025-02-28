package kinesisfirehose

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// DeliveryStream_Deserializer AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.Deserializer)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deserializer.html
type DeliveryStream_Deserializer struct {

	// HiveJsonSerDe AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deserializer.html#cfn-kinesisfirehose-deliverystream-deserializer-hivejsonserde
	HiveJsonSerDe *DeliveryStream_HiveJsonSerDe `json:"HiveJsonSerDe,omitempty"`

	// OpenXJsonSerDe AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deserializer.html#cfn-kinesisfirehose-deliverystream-deserializer-openxjsonserde
	OpenXJsonSerDe *DeliveryStream_OpenXJsonSerDe `json:"OpenXJsonSerDe,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *DeliveryStream_Deserializer) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.Deserializer"
}
