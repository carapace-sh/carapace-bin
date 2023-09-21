package conditions

import (
	"os/exec"
	"runtime"
	"slices"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/internal/condition"
)

// ConditionArch checks if the given names contain current `runtime.GOARCH`
//
//	ConditionArch("amd64")
//	ConditionArch("arm", "arm64")
func ConditionArch(s ...string) condition.Condition {
	return func(c carapace.Context) bool {
		return slices.Contains(s, runtime.GOARCH)
	}
}

// ConditionOs checks if the given names contain current `runtime.GOOS`
//
//	ConditionOs("windows")
//	ConditionOs("darwin", "linux")
func ConditionOs(s ...string) condition.Condition {
	return func(c carapace.Context) bool {
		return slices.Contains(s, runtime.GOOS)
	}
}

// ConditionPath checks if any of the given executables are in PATH
//
//	ConditionPath("git")
//	ConditionPath("carapace", "go")
func ConditionPath(s ...string) condition.Condition {
	return func(c carapace.Context) bool {
		for _, file := range s {
			if _, err := exec.LookPath(file); err == nil {
				return true
			}
		}
		return false
	}
}
