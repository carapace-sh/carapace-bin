package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_metadata_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from the branch (omit key to list all)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_metadata_help_getCmd).Standalone()

	branch_metadata_helpCmd.AddCommand(branch_metadata_help_getCmd)
}
