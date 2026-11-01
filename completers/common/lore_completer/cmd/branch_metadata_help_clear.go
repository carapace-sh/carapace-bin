package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_metadata_help_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata from the branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_metadata_help_clearCmd).Standalone()

	branch_metadata_helpCmd.AddCommand(branch_metadata_help_clearCmd)
}
