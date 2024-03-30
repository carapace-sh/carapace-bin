package transmission

import (
	"github.com/carapace-sh/carapace"
)

// ActionAbsoluteDirectories completes directories with absolute paths
func ActionAbsoluteDirectories() carapace.Action {
	return carapace.ActionDirectories().Chdir("/")
}

// ActionAbsoluteFiles completes files with absolute paths
func ActionAbsoluteFiles() carapace.Action {
	return carapace.ActionFiles().Chdir("/")
}
