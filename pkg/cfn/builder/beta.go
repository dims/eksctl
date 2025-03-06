package builder

import (
	_ "embed"
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

	customFunction := clusterTemplate.Resources["CustomEKSFunction"].(*lambda.Function)
	customFunction.Code = &lambda.Function_Code{
		ZipFile: gfnt.NewString(string(lambdaBetaPy)),
	}

	return BetaResourceRefs{
		ClusterArn: gfnt.MakeFnGetAttString("CustomEKSCluster", "PhysicalResourceId"),
	}, nil
}
