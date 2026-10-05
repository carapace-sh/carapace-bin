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
		"linux_amd64",
		"linux_386",
		"linux_arm",
		"linux_arm64",
		"darwin_amd64",
		"darwin_arm64",
		"windows_amd64",
		"windows_386",
		"windows_arm",
		"freebsd_amd64",
		"freebsd_386",
		"freebsd_arm",
		"openbsd_amd64",
		"openbsd_386",
		"openbsd_arm",
		"openbsd_arm64",
		"solaris_amd64",
	)
}
