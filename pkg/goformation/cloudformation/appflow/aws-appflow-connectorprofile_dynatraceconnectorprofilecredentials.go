package appflow

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// ConnectorProfile_DynatraceConnectorProfileCredentials AWS CloudFormation Resource (AWS::AppFlow::ConnectorProfile.DynatraceConnectorProfileCredentials)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-dynatraceconnectorprofilecredentials.html
type ConnectorProfile_DynatraceConnectorProfileCredentials struct {

	// ApiToken AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-dynatraceconnectorprofilecredentials.html#cfn-appflow-connectorprofile-dynatraceconnectorprofilecredentials-apitoken
	ApiToken *types.Value `json:"ApiToken,omitempty"`

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
func (r *ConnectorProfile_DynatraceConnectorProfileCredentials) AWSCloudFormationType() string {
	return "AWS::AppFlow::ConnectorProfile.DynatraceConnectorProfileCredentials"
}
