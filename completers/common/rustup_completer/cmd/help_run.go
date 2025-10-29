package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command with an environment configured for a given toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_runCmd).Standalone()

	helpCmd.AddCommand(help_runCmd)
}
