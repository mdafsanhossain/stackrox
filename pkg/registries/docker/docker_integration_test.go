package docker

import (
	"testing"

	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stretchr/testify/require"
)

var (
	// DefaultRegistry defaults to dockerhub
	defaultRegistry = "https://registry-1.docker.io" // variable so that it could be potentially changed
)

func TestGetMetadataIntegration(t *testing.T) {
	dockerHubClient, err := newRegistry(&v1.ImageIntegration{
		IntegrationConfig: &v1.ImageIntegration_Docker{
			Docker: &v1.DockerConfig{
				Endpoint: "https://k8s.gcr.io",
			},
		},
	})
	require.NoError(t, err)

	/*
		"k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.8"
	*/
	image := v1.Image{
		Id: "sha256:93c827f018cf3322f1ff2aa80324a0306048b0a69bc274e423071fb0d2d29d8b",
		Name: &v1.ImageName{
			Registry: "k8s.gcr.io",
			Remote:   "k8s-dns-dnsmasq-nanny-amd64",
			Tag:      "1.14.8",
		},
	}
	_, err = dockerHubClient.Metadata(&image)
	require.Nil(t, err)
}
