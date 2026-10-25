package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionConfigKeys completes keys of the project config
func ActionConfigKeys(subcommand string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		words := append([]string{"mise", "config", subcommand}, c.Args...)
		words = append(words, c.Value)
		return carapace.ActionExecCommand("mise", "__complete_word__", "--candidates", "key", "--line", strings.Join(words, " "))(func(output []byte) carapace.Action {
			return carapace.ActionValues(strings.Fields(string(output))...)
		})
	}).Tag("config keys")
}

// ActionBootstrapParts completes parts of the bootstrap process
func ActionBootstrapParts() carapace.Action {
	return carapace.ActionValues(
		"accounts",
		"compose",
		"defaults",
		"dotfiles",
		"files",
		"final-hook",
		"firewall",
		"launchd",
		"linux-systemd-units",
		"macos-defaults",
		"macos-launchd-agents",
		"mise-shell-activate",
		"packages",
		"plugins",
		"repos",
		"services",
		"shell",
		"systemd",
		"task",
		"tools",
		"user",
	)
}

// ActionShells completes shells supported by mise
func ActionShells() carapace.Action {
	return carapace.ActionValues("bash", "elvish", "fish", "nu", "powershell", "pwsh", "xonsh", "zsh")
}

// ActionCompletionShells completes shells supported by `mise completion`
func ActionCompletionShells() carapace.Action {
	return carapace.ActionValues("bash", "fish", "powershell", "pwsh", "zsh")
}

// ActionTaskCacheModes completes task cache access modes
func ActionTaskCacheModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"local-only", "Read and write only the local cache.",
		"off", "Disable task output caching for this run.",
		"read-only", "Read cached results without writing new results.",
		"read-write", "Read cached results and write new results.",
		"write-only", "Write new results without reading cached results.",
	)
}
