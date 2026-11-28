package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionVersions completes installed Ruby versions
func ActionVersions() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("rbenv", "versions", "--bare")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}).Tag("versions")
	})
}

// ActionVersionsAndSystem completes versions including the system fallback
func ActionVersionsAndSystem() carapace.Action {
	return carapace.Batch(
		carapace.ActionValues("system"),
		ActionVersions(),
	).ToA()
}

// ActionCommands completes rbenv commands
func ActionCommands() carapace.Action {
	return carapace.ActionValues(
		"commands",
		"completions",
		"exec",
		"global",
		"help",
		"hooks",
		"init",
		"local",
		"prefix",
		"rehash",
		"root",
		"shell",
		"shims",
		"version",
		"version-file",
		"version-file-read",
		"version-file-write",
		"version-name",
		"version-origin",
		"versions",
		"whence",
		"which",
	).Tag("commands")
}

// ActionShims completes shimmed commands
func ActionShims() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("rbenv", "shims", "--short")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}).Tag("shims")
	})
}

// ActionHookCommands completes commands that support hooks
func ActionHookCommands() carapace.Action {
	return carapace.ActionValues("exec", "rehash", "version-name", "version-origin", "which").Tag("commands")
}

// ActionShells completes shells supported by init
func ActionShells() carapace.Action {
	return carapace.ActionValues("bash", "fish", "ksh", "zsh")
}
