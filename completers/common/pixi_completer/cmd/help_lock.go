package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Solve environment and update the lock file without installing the environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_lockCmd).Standalone()

	helpCmd.AddCommand(help_lockCmd)
}
