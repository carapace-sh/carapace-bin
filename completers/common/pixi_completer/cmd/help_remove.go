package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes dependencies from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_removeCmd).Standalone()

	helpCmd.AddCommand(help_removeCmd)
}
