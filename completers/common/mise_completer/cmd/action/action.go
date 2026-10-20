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
		"plugins",
		"packages",
		"accounts",
		"files",
		"services",
		"firewall",
		"compose",
		"repos",
		"dotfiles",
		"mise-shell-activate",
		"macos-defaults",
		"macos-launchd-agents",
		"linux-systemd-units",
		"user",
		"tools",
		"task",
		"final-hook",
		"shell",
		"defaults",
		"launchd",
		"systemd",
	)
}

// ActionShells completes shells supported by mise
func ActionShells() carapace.Action {
	return carapace.ActionValues("bash", "elvish", "fish", "nu", "xonsh", "zsh", "pwsh", "powershell")
}

// ActionCompletionShells completes shells supported by `mise completion`
func ActionCompletionShells() carapace.Action {
	return carapace.ActionValues("bash", "fish", "powershell", "zsh", "pwsh")
}

// ActionTaskCacheModes completes task cache access modes
func ActionTaskCacheModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"read-write", "Read cached results and write new results.",
		"read-only", "Read cached results without writing new results.",
		"write-only", "Write new results without reading cached results.",
		"off", "Disable task output caching for this run.",
		"local-only", "Read and write only the local cache.",
	)
}
