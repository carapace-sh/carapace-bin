package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerControllerCmd = &cobra.Command{
	Use:     "runner-controller <command> [flags]",
	Short:   "Manage runner controllers. (EXPERIMENTAL)",
	Aliases: []string{"rc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerControllerCmd).Standalone()

	rootCmd.AddCommand(runnerControllerCmd)
}
