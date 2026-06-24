package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// TODO only relevant for darwin
var rootCmd = &cobra.Command{
	Use:   "launchctl",
	Short: "Interfaces with launchd",
	Long:  "https://ss64.com/mac/launchctl.html",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
}

func actionDomainTargets() carapace.Action {
	return carapace.ActionValuesDescribed(
		"system", "the system domain",
		"user/", "a user domain (append uid)",
		"gui/", "a GUI domain (append uid)",
		"pid/", "a process domain (append pid)",
		"login/", "a login domain (append asid)",
	)
}
