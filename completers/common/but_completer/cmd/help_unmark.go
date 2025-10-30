package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_unmarkCmd = &cobra.Command{
	Use:   "unmark",
	Short: "Removes all marks from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_unmarkCmd).Standalone()

	helpCmd.AddCommand(help_unmarkCmd)
}
