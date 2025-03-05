package builder

import (
	_ "embed"
	gfneks "github.com/weaveworks/eksctl/pkg/goformation/cloudformation/eks"

	"github.com/weaveworks/eksctl/pkg/goformation"
	gfn "github.com/weaveworks/eksctl/pkg/goformation/cloudformation"
	gfnt "github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"
)

//go:embed templates/beta-resources.yaml
var betaResourcesTemplate []byte

type BetaResourceRefs struct {
	ClusterArn *gfnt.Value
}

func AddBetaResources(clusterTemplate *gfn.Template, g *gfneks.Cluster) (BetaResourceRefs, error) {
	template, err := goformation.ParseYAML(betaResourcesTemplate)
	if err != nil {
		return BetaResourceRefs{}, err
	}
	for resourceName, resource := range template.Resources {
		clusterTemplate.Resources[resourceName] = resource
	}
	for key, output := range template.Outputs {
		clusterTemplate.Outputs[key] = output
	}
	customResource := clusterTemplate.Resources["CustomEKSCluster"].(*gfn.CustomResource)
	customResource.Properties["ResourcesVpcConfig"] = g.ResourcesVpcConfig
	customResource.Properties["Version"] = g.Version

	return BetaResourceRefs{
		ClusterArn: gfnt.MakeFnGetAttString("CustomEKSCluster", "PhysicalResourceId"),
	}, nil
}
