package kinesisanalytics

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// Application_InputProcessingConfiguration AWS CloudFormation Resource (AWS::KinesisAnalytics::Application.InputProcessingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html
type Application_InputProcessingConfiguration struct {

	// InputLambdaProcessor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html#cfn-kinesisanalytics-application-inputprocessingconfiguration-inputlambdaprocessor
	InputLambdaProcessor *Application_InputLambdaProcessor `json:"InputLambdaProcessor,omitempty"`

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
func (r *Application_InputProcessingConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.InputProcessingConfiguration"
}
