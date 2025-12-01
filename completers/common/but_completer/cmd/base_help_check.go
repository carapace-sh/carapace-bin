package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var base_help_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Fetches from the remote and checks the mergeability of the branches in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(base_help_checkCmd).Standalone()

	base_helpCmd.AddCommand(base_help_checkCmd)
}
