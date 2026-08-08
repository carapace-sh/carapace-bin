package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_scopeCmd = &cobra.Command{
	Use:   "scope <command> [flags]",
	Short: "Manage runner controller scopes. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_scopeCmd).Standalone()

	runnerControllerCmd.AddCommand(runnerController_scopeCmd)
}
