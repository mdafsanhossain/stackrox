package violationmessages

import (
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/booleanpolicy/augmentedobjs"
	"github.com/stackrox/rox/pkg/booleanpolicy/evaluator"
	"github.com/stackrox/rox/pkg/booleanpolicy/fieldnames"
	"github.com/stackrox/rox/pkg/booleanpolicy/violationmessages/printer"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/set"
)

type violationPrinter struct {
	required       set.StringSet // These fields must all be in the result, and must be valid search tags
	printerFuncKey string
}

var (
	policyFieldsToPrinters = map[storage.LifecycleStage]map[string][]violationPrinter{
		storage.LifecycleStage_DEPLOY: {
			fieldnames.AddCaps:                   {{required: set.NewStringSet(search.AddCapabilities.String()), printerFuncKey: printer.AddCapabilityKey}},
			fieldnames.CVE:                       {{required: set.NewStringSet(search.CVE.String()), printerFuncKey: printer.CveKey}},
			fieldnames.CVSS:                      {{required: set.NewStringSet(search.CVE.String()), printerFuncKey: printer.CveKey}},
			fieldnames.ContainerCPULimit:         {{required: set.NewStringSet(search.CPUCoresLimit.String()), printerFuncKey: printer.ResourceKey}},
			fieldnames.ContainerCPURequest:       {{required: set.NewStringSet(search.CPUCoresRequest.String()), printerFuncKey: printer.ResourceKey}},
			fieldnames.ContainerMemLimit:         {{required: set.NewStringSet(search.MemoryLimit.String()), printerFuncKey: printer.ResourceKey}},
			fieldnames.ContainerMemRequest:       {{required: set.NewStringSet(search.MemoryRequest.String()), printerFuncKey: printer.ResourceKey}},
			fieldnames.ContainerName:             {{required: set.NewStringSet(search.ContainerName.String()), printerFuncKey: printer.ContainerNameKey}},
			fieldnames.DisallowedAnnotation:      {{required: set.NewStringSet(search.Annotation.String()), printerFuncKey: printer.DisallowedAnnotationKey}},
			fieldnames.DisallowedImageLabel:      {{required: set.NewStringSet(search.ImageLabel.String()), printerFuncKey: printer.DisallowedImageLabelKey}},
			fieldnames.DockerfileLine:            {{required: set.NewStringSet(augmentedobjs.DockerfileLineCustomTag), printerFuncKey: printer.LineKey}},
			fieldnames.DropCaps:                  {{required: set.NewStringSet(search.DropCapabilities.String()), printerFuncKey: printer.DropCapabilityKey}},
			fieldnames.EnvironmentVariable:       {{required: set.NewStringSet(augmentedobjs.EnvironmentVarCustomTag), printerFuncKey: printer.EnvKey}},
			fieldnames.ExposedPort:               {{required: set.NewStringSet(search.Port.String()), printerFuncKey: printer.PortKey}},
			fieldnames.FixedBy:                   {{required: set.NewStringSet(search.CVE.String()), printerFuncKey: printer.CveKey}},
			fieldnames.ImageAge:                  {{required: set.NewStringSet(search.ImageCreatedTime.String()), printerFuncKey: printer.ImageAgeKey}},
			fieldnames.ImageComponent:            {{required: set.NewStringSet(augmentedobjs.ComponentAndVersionCustomTag), printerFuncKey: printer.ComponentKey}},
			fieldnames.ImageOS:                   {{required: set.NewStringSet(search.ImageOS.String()), printerFuncKey: printer.ImageOSKey}},
			fieldnames.ImageRegistry:             {{required: set.StringSet{}, printerFuncKey: printer.ImageDetailsKey}},
			fieldnames.ImageRemote:               {{required: set.StringSet{}, printerFuncKey: printer.ImageDetailsKey}},
			fieldnames.ImageScanAge:              {{required: set.NewStringSet(search.ImageScanTime.String()), printerFuncKey: printer.ImageScanAgeKey}},
			fieldnames.ImageTag:                  {{required: set.StringSet{}, printerFuncKey: printer.ImageDetailsKey}},
			fieldnames.ImageUser:                 {{required: set.StringSet{}, printerFuncKey: printer.ImageUserKey}},
			fieldnames.MinimumRBACPermissions:    {{required: set.NewStringSet(search.ServiceAccountPermissionLevel.String()), printerFuncKey: printer.RbacKey}},
			fieldnames.Namespace:                 {{required: set.NewStringSet(search.Namespace.String()), printerFuncKey: printer.NamespaceKey}},
			fieldnames.PortExposure:              {{required: set.NewStringSet(search.ExposureLevel.String()), printerFuncKey: printer.PortExposureKey}},
			fieldnames.PrivilegedContainer:       {{required: set.NewStringSet(search.Privileged.String()), printerFuncKey: printer.PrivilegedKey}},
			fieldnames.ExposedPortProtocol:       {{required: set.NewStringSet(search.Port.String()), printerFuncKey: printer.PortKey}},
			fieldnames.ReadOnlyRootFS:            {{required: set.NewStringSet(search.ReadOnlyRootFilesystem.String()), printerFuncKey: printer.ReadOnlyRootFSKey}},
			fieldnames.RequiredAnnotation:        {{required: set.NewStringSet(search.Annotation.String()), printerFuncKey: printer.RequiredAnnotationKey}},
			fieldnames.RequiredImageLabel:        {{required: set.NewStringSet(search.ImageLabel.String()), printerFuncKey: printer.RequiredImageLabelKey}},
			fieldnames.RequiredLabel:             {{required: set.NewStringSet(search.Label.String()), printerFuncKey: printer.RequiredLabelKey}},
			fieldnames.ServiceAccount:            {{required: set.NewStringSet(search.ServiceAccountName.String()), printerFuncKey: printer.ServiceAccountKey}},
			fieldnames.UnexpectedProcessExecuted: {{required: set.NewStringSet(augmentedobjs.NotInBaselineCustomTag), printerFuncKey: printer.ProcessBaselineKey}},
			fieldnames.UnscannedImage:            {{required: set.NewStringSet(augmentedobjs.ImageScanCustomTag), printerFuncKey: printer.ImageScanKey}},
			fieldnames.VolumeDestination:         {{required: set.NewStringSet(search.VolumeName.String()), printerFuncKey: printer.VolumeKey}},
			fieldnames.VolumeName:                {{required: set.NewStringSet(search.VolumeName.String()), printerFuncKey: printer.VolumeKey}},
			fieldnames.VolumeSource:              {{required: set.NewStringSet(search.VolumeName.String()), printerFuncKey: printer.VolumeKey}},
			fieldnames.VolumeType:                {{required: set.NewStringSet(search.VolumeName.String()), printerFuncKey: printer.VolumeKey}},
			fieldnames.WritableHostMount:         {{required: set.NewStringSet(search.VolumeName.String()), printerFuncKey: printer.VolumeKey}},
			fieldnames.WritableMountedVolume:     {{required: set.NewStringSet(search.VolumeName.String()), printerFuncKey: printer.VolumeKey}},
		},
		storage.LifecycleStage_BUILD: {
			fieldnames.CVE:                  {{required: set.NewStringSet(search.CVE.String()), printerFuncKey: printer.CveKey}},
			fieldnames.CVSS:                 {{required: set.NewStringSet(search.CVE.String()), printerFuncKey: printer.CveKey}},
			fieldnames.DisallowedImageLabel: {{required: set.NewStringSet(search.ImageLabel.String()), printerFuncKey: printer.DisallowedImageLabelKey}},
			fieldnames.DockerfileLine:       {{required: set.NewStringSet(augmentedobjs.DockerfileLineCustomTag), printerFuncKey: printer.LineKey}},
			fieldnames.FixedBy:              {{required: set.NewStringSet(search.CVE.String()), printerFuncKey: printer.CveKey}},
			fieldnames.ImageAge:             {{required: set.NewStringSet(search.ImageCreatedTime.String()), printerFuncKey: printer.ImageAgeKey}},
			fieldnames.ImageComponent:       {{required: set.NewStringSet(augmentedobjs.ComponentAndVersionCustomTag), printerFuncKey: printer.ComponentKey}},
			fieldnames.ImageOS:              {{required: set.NewStringSet(search.ImageOS.String()), printerFuncKey: printer.ImageOSKey}},
			fieldnames.ImageRegistry:        {{required: set.StringSet{}, printerFuncKey: printer.ImageDetailsKey}},
			fieldnames.ImageRemote:          {{required: set.StringSet{}, printerFuncKey: printer.ImageDetailsKey}},
			fieldnames.ImageScanAge:         {{required: set.NewStringSet(search.ImageScanTime.String()), printerFuncKey: printer.ImageScanAgeKey}},
			fieldnames.ImageTag:             {{required: set.StringSet{}, printerFuncKey: printer.ImageDetailsKey}},
			fieldnames.ImageUser:            {{required: set.StringSet{}, printerFuncKey: printer.ImageUserKey}},
			fieldnames.RequiredImageLabel:   {{required: set.NewStringSet(search.ImageLabel.String()), printerFuncKey: printer.RequiredImageLabelKey}},
			fieldnames.UnscannedImage:       {{required: set.NewStringSet(augmentedobjs.ImageScanCustomTag), printerFuncKey: printer.ImageScanKey}},
		},
	}

	// runtime policy fields
	requiredProcessFields = set.NewFrozenStringSet(search.ProcessName.String(), search.ProcessAncestor.String(),
		search.ProcessUID.String(), search.ProcessArguments.String(), augmentedobjs.NotInBaselineCustomTag)
	requiredKubeEventFields = set.NewFrozenStringSet(augmentedobjs.KubernetesAPIVerbCustomTag, augmentedobjs.KubernetesResourceCustomTag)
)

func containsAllRequiredFields(fieldMap map[string][]string, required set.StringSet) bool {
	for field := range required {
		if _, ok := fieldMap[field]; !ok {
			return false
		}
	}
	return true
}

func lookupViolationPrinters(stage storage.LifecycleStage, section *storage.PolicySection, fieldMap map[string][]string) []printer.Func {
	keys := set.NewStringSet()
	if printersAndFields, ok := policyFieldsToPrinters[stage]; ok {
		for _, group := range section.GetPolicyGroups() {
			if printerMD, ok := printersAndFields[group.GetFieldName()]; ok {
				for _, p := range printerMD {
					if containsAllRequiredFields(fieldMap, p.required) {
						keys.Add(p.printerFuncKey)
					}
				}
			}
		}
	}
	if len(keys) == 0 {
		return nil
	}
	return printer.GetFuncs(keys)
}

func checkForProcessViolation(result *evaluator.Result) bool {
	for _, fieldMap := range result.Matches {
		for k := range fieldMap {
			if requiredProcessFields.Contains(k) {
				return true
			}
		}
	}
	return false
}

func checkForKubeEventViolation(result *evaluator.Result) bool {
	for _, fieldMap := range result.Matches {
		for k := range fieldMap {
			if requiredKubeEventFields.Contains(k) {
				return true
			}
		}
	}
	return false
}

// Render creates violation messages based on evaluation results
func Render(
	stage storage.LifecycleStage,
	section *storage.PolicySection,
	result *evaluator.Result,
	indicator *storage.ProcessIndicator,
	kubeEvent *storage.KubernetesEvent,
) ([]*storage.Alert_Violation, bool, bool, error) {
	errorList := errorhelpers.NewErrorList("violation printer")
	messages := set.NewStringSet()
	for _, fieldMap := range result.Matches {
		printers := lookupViolationPrinters(stage, section, fieldMap)
		if len(printers) == 0 {
			continue
		}
		for _, printerFunc := range printers {
			messagesForResult, err := printerFunc(fieldMap)
			if err != nil {
				errorList.AddError(err)
				continue
			}
			messages.AddAll(messagesForResult...)
		}
	}

	isProcessViolation := indicator != nil && checkForProcessViolation(result)
	isKubeEventViolation := kubeEvent != nil && checkForKubeEventViolation(result)
	if len(messages) == 0 && !isProcessViolation && !isKubeEventViolation {
		errorList.AddError(errors.New("missing messages"))
	}

	alertViolations := make([]*storage.Alert_Violation, 0, len(messages))
	// Sort messages for consistency in output. This is important because we
	// depend on these messages being equal when deduping updates to alerts.
	for _, message := range messages.AsSortedSlice(func(i, j string) bool {
		return i < j
	}) {
		alertViolations = append(alertViolations, &storage.Alert_Violation{Message: message})
	}
	return alertViolations, isProcessViolation, isKubeEventViolation, errorList.ToError()
}
