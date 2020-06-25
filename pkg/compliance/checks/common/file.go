package common

import (
	"fmt"
	"path/filepath"

	"github.com/stackrox/rox/generated/internalapi/compliance"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/compliance/checks/standards"
)

// SystemdOwnershipCheck checks the users and groups of the file
func SystemdOwnershipCheck(file, user, group string) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          systemdOwnershipCheckFunc(file, user, group, false),
		InterpretationText: fmt.Sprintf("StackRox checks that the systemd file %s on each node is owned by user %q and group %q", file, user, group),
	}
}

// OwnershipCheck checks the users and groups of the file
func OwnershipCheck(file, user, group string) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          ownershipCheckFunc(file, user, group, false),
		InterpretationText: fmt.Sprintf("StackRox checks that the file %s on each node is owned by user %q and group %q", file, user, group),
	}
}

// OptionalOwnershipCheck checks the users and groups of the file if it exists. If it does not exist, then the check passes
func OptionalOwnershipCheck(file, user, group string) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          ownershipCheckFunc(file, user, group, true),
		InterpretationText: fmt.Sprintf("StackRox checks that the file %s on each node (if existing) is owned by user %q and group %q", file, user, group),
	}
}

// RecursiveOwnershipCheck is a framework Check for recursively checking the ownership
func RecursiveOwnershipCheck(dir, user, group string) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          recursiveOwnershipCheckFunc(dir, user, group, false),
		InterpretationText: fmt.Sprintf("StackRox checks that all files under the path %s are owned by user %q and group %q", dir, user, group),
	}
}

// RecursiveOwnershipCheckIfDirExists is a framework Check for recursively checking the ownership
func RecursiveOwnershipCheckIfDirExists(dir, user, group string) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          recursiveOwnershipCheckFunc(dir, user, group, true),
		InterpretationText: fmt.Sprintf("StackRox checks that all files under the path %s are owned by user %q and group %q", dir, user, group),
	}
}

// CheckRecursiveOwnership checks the files against the passed user and group
func CheckRecursiveOwnership(f *compliance.File, user, group string) []*storage.ComplianceResultValue_Evidence {
	var results []*storage.ComplianceResultValue_Evidence
	results = append(results, ownershipCheck(f, user, group)...)
	for _, f := range f.Children {
		results = append(results, CheckRecursiveOwnership(f, user, group)...)
	}
	return results
}

func ownershipCheckFunc(path, user, group string, optional bool) standards.Check {
	return func(complianceData *standards.ComplianceData) []*storage.ComplianceResultValue_Evidence {
		f, ok := complianceData.Files[path]
		if !ok && optional {
			return PassListf("File %q does not exist on host, therefore check is not applicable", path)
		} else if !ok {
			return FailListf("File %q could not be found in scraped data", path)
		}
		return ownershipCheck(f, user, group)
	}
}

func ownershipCheck(f *compliance.File, user, group string) []*storage.ComplianceResultValue_Evidence {
	var results []*storage.ComplianceResultValue_Evidence
	var fail bool
	if !HasOwnershipUser(f, user) {
		fail = true
		results = append(results, Failf("Expected user %q on file %q, but found %q", user, f.GetPath(), f.GetUserName()))
	}
	if !HasOwnershipGroup(f, group) {
		fail = true
		results = append(results, Failf("Expected group %q on file %q, but found %q", group, f.GetPath(), f.GetGroupName()))
	}
	if !fail {
		results = append(results, Passf("Found group %q and user %q on file %q", group, user, f.GetPath()))
	}
	return results
}

// PermissionCheck checks the permissions of the file
func PermissionCheck(file string, permissions uint32) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          permissionCheckFunc(file, permissions, false),
		InterpretationText: fmt.Sprintf("StackRox checks that the permissions on file %s on each node are set to '%#o'", file, permissions),
	}
}

// OptionalSystemdOwnershipCheck checks the users and groups of the file if it exists. If it does not exist, then the check passes
func OptionalSystemdOwnershipCheck(file, user, group string) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          systemdOwnershipCheckFunc(file, user, group, true),
		InterpretationText: fmt.Sprintf("StackRox checks that the systemd file %s on each node is owned by user %q and group %q", file, user, group),
	}
}

// SystemdPermissionCheck checks the permissions of the file
func SystemdPermissionCheck(file string, permissions uint32) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          systemdPermissionCheckFunc(file, permissions, false),
		InterpretationText: fmt.Sprintf("StackRox checks that the permissions on the systemd file %s on each node are set to '%#o'", file, permissions),
	}
}

// HasOwnershipUser checks the user owner on a file
func HasOwnershipUser(f *compliance.File, user string) bool {
	return f.GetUserName() == user
}

// HasOwnershipGroup checks the group owner on a file
func HasOwnershipGroup(f *compliance.File, group string) bool {
	return f.GetGroupName() == group
}

func permissionCheck(f *compliance.File, permissions uint32) (*storage.ComplianceResultValue_Evidence, bool) {
	if !HasPermissions(f, permissions) {
		return Failf("Expected permissions '%#o' on file %q, but found '%#o'", permissions, f.GetPath(), f.GetPermissions()), true
	}
	return Passf("Found permissions '%#o' on file %q", permissions, f.GetPath()), false
}

// HasPermissions checks the permissions on a file
func HasPermissions(f *compliance.File, permissionLevel uint32) bool {
	return f.GetPermissions() == permissionLevel || f.GetPermissions() < permissionLevel
}

// CheckRecursivePermissions does the actual checking of the files
func CheckRecursivePermissions(f *compliance.File, permissions uint32) ([]*storage.ComplianceResultValue_Evidence, bool) {
	var results []*storage.ComplianceResultValue_Evidence
	result, stopNow := permissionCheck(f, permissions)
	results = append(results, result)
	if stopNow {
		return results, stopNow
	}
	for _, child := range f.Children {
		result, stopNow := CheckRecursivePermissions(child, permissions)
		results = append(results, result...)
		if stopNow {
			return results, stopNow
		}
	}
	return results, false
}

func permissionCheckFunc(path string, permissions uint32, optional bool) standards.Check {
	return func(complianceData *standards.ComplianceData) []*storage.ComplianceResultValue_Evidence {
		f, ok := complianceData.Files[path]
		if !ok && optional {
			return PassListf("File %q does not exist on host, therefore check is not applicable", path)
		} else if !ok {
			return FailListf("File %q could not be found in scraped data", path)
		}
		result, _ := permissionCheck(f, permissions)
		return []*storage.ComplianceResultValue_Evidence{result}
	}
}

// OptionalPermissionCheck checks the permissions of the optional file
func OptionalPermissionCheck(file string, permissions uint32) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          permissionCheckFunc(file, permissions, true),
		InterpretationText: fmt.Sprintf("StackRox checks that the permissions on file %s on each node (if existing) are set to '%#o'", file, permissions),
	}
}

// OptionalSystemdPermissionCheck checks the permissions of the file
func OptionalSystemdPermissionCheck(file string, permissions uint32) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          systemdPermissionCheckFunc(file, permissions, true),
		InterpretationText: fmt.Sprintf("StackRox checks that the permissions on the systemd file %s on each node are set to '%#o'", file, permissions),
	}
}

func recursivePermissionCheckFunc(path string, permissions uint32) standards.Check {
	return func(complianceData *standards.ComplianceData) []*storage.ComplianceResultValue_Evidence {
		f, ok := complianceData.Files[path]
		if !ok {
			return FailListf("File %q could not be found in scraped data", path)
		}
		results, _ := CheckRecursivePermissions(f, permissions)
		return results
	}
}

// RecursivePermissionCheckWithFileExtIfDirExists recursively checks the permissions of the file with given extension
func RecursivePermissionCheckWithFileExtIfDirExists(dir, ext string, permissions uint32) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          recursivePermissionCheckWithFileExtFunc(dir, ext, permissions, true),
		InterpretationText: fmt.Sprintf("StackRox checks that the permissions of files with extension %s under the path %s on each node are set to '%#o'", ext, dir, permissions),
	}
}

func recursivePermissionCheckWithFileExtFunc(path, fileExtension string, permissions uint32, optional bool) standards.Check {
	return func(complianceData *standards.ComplianceData) []*storage.ComplianceResultValue_Evidence {
		f, ok := complianceData.Files[path]
		if !ok && optional {
			return PassListf("File %q does not exist on host, therefore check is not applicable", path)
		} else if !ok {
			return FailListf("File %q could not be found in scraped data", path)
		}
		results, _ := CheckRecursivePermissionWithFileExt(f, fileExtension, permissions)
		return results
	}
}

// CheckRecursivePermissionWithFileExt does the actual checking of the files
func CheckRecursivePermissionWithFileExt(f *compliance.File, fileExtension string, permissions uint32) ([]*storage.ComplianceResultValue_Evidence, bool) {
	if filepath.Ext(f.GetPath()) == fileExtension {
		result, stopNow := permissionCheck(f, permissions)
		return []*storage.ComplianceResultValue_Evidence{result}, stopNow
	}
	var results []*storage.ComplianceResultValue_Evidence
	for _, child := range f.Children {
		childResults, failNow := CheckRecursivePermissionWithFileExt(child, fileExtension, permissions)
		results = append(results, childResults...)
		if failNow {
			return results, true
		}
	}
	return results, false
}

func recursiveOwnershipCheckFunc(path, user, group string, optional bool) standards.Check {
	return func(complianceData *standards.ComplianceData) []*storage.ComplianceResultValue_Evidence {
		f, ok := complianceData.Files[path]
		if !ok && optional {
			return PassListf("File %q does not exist on host, therefore check is not applicable", path)
		} else if !ok {
			return FailListf("File %q could not be found in scraped data", path)
		}
		return CheckRecursiveOwnership(f, user, group)
	}
}

func systemdOwnershipCheckFunc(path, user, group string, optional bool) standards.Check {
	return func(complianceData *standards.ComplianceData) []*storage.ComplianceResultValue_Evidence {
		f, ok := complianceData.SystemdFiles[path]
		if !ok {
			if optional {
				return PassListf("Service %q does not exist on host, therefore check is not applicable", path)
			}
			return FailListf("Service %q could not be found in scraped data", path)
		}
		return ownershipCheck(f, user, group)
	}
}

// RecursivePermissionCheck recursively checks the permissions of the file
func RecursivePermissionCheck(file string, permissions uint32) *standards.CheckAndInterpretation {
	return &standards.CheckAndInterpretation{
		CheckFunc:          recursivePermissionCheckFunc(file, permissions),
		InterpretationText: fmt.Sprintf("StackRox checks that the permissions of all files under the path %s on each node are set to '%#o'", file, permissions),
	}
}

func systemdPermissionCheckFunc(path string, permissions uint32, optional bool) standards.Check {
	return func(complianceData *standards.ComplianceData) []*storage.ComplianceResultValue_Evidence {
		f, ok := complianceData.SystemdFiles[path]
		if !ok {
			if optional {
				return PassListf("Service %q does not exist on host, therefore check is not applicable", path)
			}
			return FailListf("Service %q could not be found in scraped data", path)
		}
		result, _ := permissionCheck(f, permissions)
		return []*storage.ComplianceResultValue_Evidence{result}
	}
}
