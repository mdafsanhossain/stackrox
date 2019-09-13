package docker

import (
	"github.com/stackrox/rox/central/compliance/checks/common"
	"github.com/stackrox/rox/central/compliance/framework"
	"github.com/stackrox/rox/pkg/docker/types"
	"github.com/stackrox/rox/pkg/set"
)

func init() {
	framework.MustRegisterChecks(
		framework.NewCheckFromFunc(
			framework.CheckMetadata{
				ID:                 "CIS_Docker_v1_2_0:6_1",
				Scope:              framework.NodeKind,
				InterpretationText: "StackRox checks how many of the images present on each node are actually in use",
				DataDependencies:   []string{"HostScraped"},
			}, imageSprawl),
		framework.NewCheckFromFunc(
			framework.CheckMetadata{
				ID:                 "CIS_Docker_v1_2_0:6_2",
				Scope:              framework.NodeKind,
				InterpretationText: "StackRox checks how many of the containers present on each node are actually running",
				DataDependencies:   []string{"HostScraped"},
			}, containerSprawl),
	)
}

func imageSprawl(funcCtx framework.ComplianceContext) {
	common.PerNodeCheckWithDockerData(func(ctx framework.ComplianceContext, data *types.Data) {
		imageSet := set.NewStringSet()
		for _, c := range data.Containers {
			imageSet.Add(c.Image)
		}
		framework.Notef(ctx, "There are %d images in use out of %d", imageSet.Cardinality(), len(data.Images))
	})(funcCtx)
}

func containerSprawl(funcCtx framework.ComplianceContext) {
	common.PerNodeCheckWithDockerData(func(ctx framework.ComplianceContext, data *types.Data) {
		var runningContainers int
		for _, c := range data.Containers {
			if c.State != nil && c.State.Running {
				runningContainers++
			}
		}
		framework.Notef(ctx, "There are %d containers in use out of %d", runningContainers, len(data.Containers))
	})(funcCtx)
}
