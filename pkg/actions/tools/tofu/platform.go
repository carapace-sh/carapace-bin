package tofu

import (
	"github.com/carapace-sh/carapace"
)

// ActionPlatforms completes target platforms for provider packages
//
//	linux_amd64
//	darwin_arm64
func ActionPlatforms() carapace.Action {
	return carapace.ActionValues(
		"darwin_amd64",
		"darwin_arm64",
		"freebsd_386",
		"freebsd_amd64",
		"freebsd_arm",
		"linux_386",
		"linux_amd64",
		"linux_arm",
		"linux_arm64",
		"openbsd_386",
		"openbsd_amd64",
		"openbsd_arm",
		"openbsd_arm64",
		"solaris_amd64",
		"windows_386",
		"windows_amd64",
		"windows_arm",
		"windows_arm64",
	)
}
