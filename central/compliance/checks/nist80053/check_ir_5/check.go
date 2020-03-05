package checkir5

import (
	"github.com/stackrox/rox/central/compliance/checks/common"
	"github.com/stackrox/rox/central/compliance/framework"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
)

const (
	controlID = "NIST_SP_800_53_Rev_4:IR_5"

	phase = storage.LifecycleStage_RUNTIME
)

var (
	interpretationText = `This control requires tracking and documenting information security incidents.

For this control, ` + common.AnyPolicyInLifeCycleInterpretation(phase)
)

func init() {
	framework.MustRegisterNewCheckIfFlagEnabled(
		framework.CheckMetadata{
			ID:                 controlID,
			Scope:              framework.ClusterKind,
			DataDependencies:   []string{"Policies"},
			InterpretationText: interpretationText,
		},
		func(ctx framework.ComplianceContext) {
			framework.Pass(ctx, "The StackRox Kubernetes Security Platform is installed and tracking information security incidents.")
			common.CheckAnyPolicyInLifeCycle(ctx, phase)
		}, features.NistSP800_53)
}
