package builder

import (
	_ "embed"
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/cloudformation"
	gfneks "github.com/weaveworks/eksctl/pkg/goformation/cloudformation/eks"
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/lambda"

	"github.com/weaveworks/eksctl/pkg/goformation"
	gfn "github.com/weaveworks/eksctl/pkg/goformation/cloudformation"
	gfnt "github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"
)

//go:embed templates/beta-resources.yaml
var betaResourcesTemplate []byte

//go:embed templates/beta.py
var lambdaBetaPy []byte

func addBetaResources(clusterName string, clusterTemplate *gfn.Template, g *gfneks.Cluster, roleArn, iamARN string) error {
	template, err := goformation.ParseYAML(betaResourcesTemplate)
	if err != nil {
		return err
	}
	for resourceName, resource := range template.Resources {
		clusterTemplate.Resources[resourceName] = resource
	}
	for key, output := range template.Outputs {
		clusterTemplate.Outputs[key] = output
	}
	customResource := clusterTemplate.Resources["ControlPlane"].(*gfn.CustomResource)
	if g.AccessConfig != nil {
		customResource.Properties["AccessConfig"] = g.AccessConfig
	}
	if g.BootstrapSelfManagedAddons != nil {
		customResource.Properties["BootstrapSelfManagedAddons"] = g.BootstrapSelfManagedAddons
	}
	if g.ComputeConfig != nil {
		customResource.Properties["ComputeConfig"] = g.ComputeConfig
	}
	if g.EncryptionConfig != nil {
		customResource.Properties["EncryptionConfig"] = g.EncryptionConfig
	}
	if g.KubernetesNetworkConfig != nil {
		customResource.Properties["KubernetesNetworkConfig"] = g.KubernetesNetworkConfig
	}
	if g.Logging != nil {
		customResource.Properties["Logging"] = g.Logging
	}
	if g.Name != nil {
		customResource.Properties["Name"] = g.Name
	}
	if g.OutpostConfig != nil {
		customResource.Properties["OutpostConfig"] = g.OutpostConfig
	}
	if g.RemoteNetworkConfig != nil {
		customResource.Properties["RemoteNetworkConfig"] = g.RemoteNetworkConfig
	}
	if g.ResourcesVpcConfig != nil {
		customResource.Properties["ResourcesVpcConfig"] = g.ResourcesVpcConfig
	}
	if g.RoleArn != nil {
		customResource.Properties["RoleArn"] = g.RoleArn
	}
	if g.StorageConfig != nil {
		customResource.Properties["StorageConfig"] = g.StorageConfig
	}
	if g.Tags != nil {
		g.Tags = append(g.Tags, cloudformation.Tag{
			Key:   gfnt.NewString("Name"),
			Value: gfnt.NewString(clusterName + "/ControlPlane"),
		})
		customResource.Properties["Tags"] = g.Tags
	} else {
		customResource.Properties["Tags"] = []cloudformation.Tag{
			{
				Key:   gfnt.NewString("Name"),
				Value: gfnt.NewString(clusterName + "/ControlPlane"),
			},
		}
	}
	if g.UpgradePolicy != nil {
		customResource.Properties["UpgradePolicy"] = g.UpgradePolicy
	}
	if g.Version != nil {
		customResource.Properties["Version"] = g.Version
	}
	if g.ZonalShiftConfig != nil {
		customResource.Properties["ZonalShiftConfig"] = g.ZonalShiftConfig
	}

	customResource.Properties["IAMPrincipalArn"] = gfnt.NewString(iamARN)
	customResource.Properties["STSRoleArn"] = gfnt.NewString(roleArn)

	customFunction := clusterTemplate.Resources["CustomEKSFunction"].(*lambda.Function)
	customFunction.Code = &lambda.Function_Code{
		ZipFile: gfnt.NewString(string(lambdaBetaPy)),
	}

	return nil
}
