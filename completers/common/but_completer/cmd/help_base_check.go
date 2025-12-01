package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_base_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Fetches from the remote and checks the mergeability of the branches in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_base_checkCmd).Standalone()

	help_baseCmd.AddCommand(help_base_checkCmd)
}
