package tests

import (
	"context"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/clientconn"
	"bitbucket.org/stack-rox/apollo/pkg/images"
	"bitbucket.org/stack-rox/apollo/pkg/protoconv"
	"bitbucket.org/stack-rox/apollo/pkg/search"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClusters(t *testing.T) {

	conn, err := clientconn.UnauthenticatedGRPCConnection(apiEndpoint)
	require.NoError(t, err)

	service := v1.NewClustersServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	clusters, err := service.GetClusters(ctx, &empty.Empty{})
	cancel()
	require.NoError(t, err)
	require.Len(t, clusters.GetClusters(), 1)

	c := clusters.GetClusters()[0]
	assert.Equal(t, v1.ClusterType_KUBERNETES_CLUSTER, c.GetType())
	assert.Equal(t, `remote`, c.GetName())

	img := images.GenerateImageFromString(c.GetPreventImage())
	assert.Equal(t, `stackrox/prevent`, img.GetName().GetRemote())
	if sha, ok := os.LookupEnv(`PREVENT_IMAGE_TAG`); ok {
		assert.Equal(t, sha, img.GetName().GetTag())
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
	cByID, err := service.GetCluster(ctx, &v1.ResourceByID{Id: c.GetId()})
	cancel()
	require.NoError(t, err)

	cByID.GetCluster().LastContact = c.GetLastContact()
	assert.Equal(t, c, cByID.GetCluster())
	for _, f := range cByID.GetFiles() {
		assert.NotEmpty(t, f.GetName())
		assert.NotEmpty(t, f.GetContent())
		if strings.HasSuffix(f.GetName(), ".sh") {
			assert.True(t, f.GetExecutable())
		}
	}
}

func TestDeployments(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := clientconn.UnauthenticatedGRPCConnection(apiEndpoint)
	require.NoError(t, err)

	service := v1.NewDeploymentServiceClient(conn)

	qb := search.NewQueryBuilder().AddStrings(search.DeploymentName, "central").AddStrings(search.DeploymentName, "sensor")
	deployments, err := service.ListDeployments(ctx, &v1.RawQuery{
		Query: qb.Query(),
	})
	require.NoError(t, err)
	require.Len(t, deployments.GetDeployments(), 2)

	var centralDeployment, sensorDeployment *v1.Deployment

	for _, d := range deployments.GetDeployments() {
		if d.GetName() == `central` {
			centralDeployment, err = retrieveDeployment(service, d)
			require.NoError(t, err)
		} else if d.GetName() == `sensor` {
			sensorDeployment, err = retrieveDeployment(service, d)
			require.NoError(t, err)
		}
	}

	require.NotNil(t, centralDeployment)
	require.NotNil(t, sensorDeployment)

	verifyCentralDeployment(t, centralDeployment)
	verifySensorDeployment(t, sensorDeployment)

	centralByID, err := service.GetDeployment(ctx, &v1.ResourceByID{Id: centralDeployment.GetId()})
	require.NoError(t, err)
	assert.Equal(t, centralDeployment, centralByID)

	sensorByID, err := service.GetDeployment(ctx, &v1.ResourceByID{Id: sensorDeployment.GetId()})
	require.NoError(t, err)
	assert.Equal(t, sensorDeployment, sensorByID)
}

func verifyCentralDeployment(t *testing.T, centralDeployment *v1.Deployment) {
	verifyDeployment(t, centralDeployment)
	assert.Equal(t, "central", protoconv.ConvertDeploymentKeyValues(centralDeployment.GetLabels())["app"])

	require.Len(t, centralDeployment.GetContainers(), 1)
	c := centralDeployment.GetContainers()[0]

	assert.Equal(t, `stackrox/prevent`, c.GetImage().GetName().GetRemote())
	if sha, ok := os.LookupEnv(`PREVENT_IMAGE_TAG`); ok {
		assert.Equal(t, sha, c.GetImage().GetName().GetTag())
	}

	require.Len(t, c.GetSecrets(), 2)
	paths := make([]string, 0, 2)
	for _, secret := range c.GetSecrets() {
		paths = append(paths, secret.GetPath())
	}
	sort.Slice(paths, func(i, j int) bool {
		return paths[i] < paths[j]
	})
	expectedPathPrefixes := []string{
		"/run/secrets/stackrox.io/certs",
		"/run/secrets/stackrox.io/jwt",
	}
	for i, path := range paths {
		assert.True(t, strings.HasPrefix(path, expectedPathPrefixes[i]))
	}

	require.Len(t, c.GetPorts(), 1)
	p := c.GetPorts()[0]
	assert.Equal(t, int32(443), p.GetContainerPort())
	assert.Equal(t, "TCP", p.GetProtocol())
}

func verifySensorDeployment(t *testing.T, sensorDeployment *v1.Deployment) {
	verifyDeployment(t, sensorDeployment)
	assert.Equal(t, "sensor", protoconv.ConvertDeploymentKeyValues(sensorDeployment.GetLabels())["app"])

	require.Len(t, sensorDeployment.GetContainers(), 1)
	c := sensorDeployment.GetContainers()[0]

	assert.Equal(t, `stackrox/prevent`, c.GetImage().GetName().GetRemote())
	if sha, ok := os.LookupEnv(`PREVENT_IMAGE_TAG`); ok {
		assert.Equal(t, sha, c.GetImage().GetName().GetTag())
	}

	require.Len(t, c.GetSecrets(), 1)
	s := c.GetSecrets()[0]
	assert.True(t, strings.HasPrefix(s.GetPath(), "/run/secrets/stackrox.io/"))
}

func verifyDeployment(t *testing.T, deployment *v1.Deployment) {
	assert.Equal(t, "Deployment", deployment.GetType())
	assert.Equal(t, int64(1), deployment.GetReplicas())
	assert.NotEmpty(t, deployment.GetId())
	assert.NotEmpty(t, deployment.GetVersion())
	assert.NotEmpty(t, deployment.GetUpdatedAt())
}

func TestImages(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := clientconn.UnauthenticatedGRPCConnection(apiEndpoint)
	require.NoError(t, err)

	service := v1.NewImageServiceClient(conn)

	images, err := service.ListImages(ctx, &v1.RawQuery{})
	require.NoError(t, err)

	require.NotEmpty(t, images.GetImages())

	imageMap := make(map[string][]*v1.Image)
	for _, img := range images.GetImages() {
		image, err := service.GetImage(ctx, &v1.ResourceByID{Id: img.GetSha()})
		assert.NoError(t, err)
		imageMap[image.GetName().GetRegistry()] = append(imageMap[image.GetName().GetRegistry()], image)
	}

	const dockerRegistry = `docker.io`

	require.NotEmpty(t, imageMap[dockerRegistry])

	foundPreventImage := false

	for _, img := range imageMap[dockerRegistry] {
		if img.GetName().GetRemote() == `stackrox/prevent` {
			foundPreventImage = true

			if sha, ok := os.LookupEnv(`PREVENT_IMAGE_TAG`); ok {
				assert.Equal(t, sha, img.GetName().GetTag())
			}
		}
	}

	assert.True(t, foundPreventImage)
}
