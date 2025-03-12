package v1alpha5

import "os"

func (c *ClusterConfig) IsEksBetaEndpoint() bool {
	eksEndpoint := os.Getenv("AWS_EKS_ENDPOINT")
	if eksEndpoint == "" {
		eksEndpoint = os.Getenv("AWS_ENDPOINT_URL_EKS")
	}
	return eksEndpoint == "https://api.beta.us-west-2.wesley.amazonaws.com"
}
