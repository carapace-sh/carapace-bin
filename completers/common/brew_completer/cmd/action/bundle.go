package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var bundleTypes = []string{"cargo", "cask", "flatpak", "formula", "go", "krew", "mas", "npm", "tap", "uv", "vscode", "winget"}

// ActionBundleEntries completes entries of the Brewfile
func ActionBundleEntries(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"bundle", "list"}
		if flag := cmd.Flag("file"); flag != nil && flag.Changed {
			args = append(args, "--file", flag.Value.String())
		}
		if flag := cmd.Flag("global"); flag != nil && flag.Changed {
			args = append(args, "--global")
		}
		for _, t := range bundleTypes {
			if flag := cmd.Flag(t); flag != nil && flag.Changed {
				args = append(args, "--"+t)
			}
		}

		c.Setenv("HOMEBREW_NO_AUTO_UPDATE", "1")
		return carapace.ActionExecCommand("brew", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}).Tag("bundle entries")
	})
}

// ActionBundlePackages completes packages of the Brewfile entry type selected by flag
func ActionBundlePackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if flag := cmd.Flag("tap"); flag != nil && flag.Changed {
			return brew.ActionInstalledTaps()
		}

		for _, t := range []string{"cargo", "flatpak", "go", "krew", "npm", "uv", "vscode"} {
			if flag := cmd.Flag(t); flag != nil && flag.Changed {
				return carapace.ActionValues()
			}
		}
		return ActionSearch(cmd)
	})
}
