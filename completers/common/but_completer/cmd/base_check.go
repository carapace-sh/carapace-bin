package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var base_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Fetches remotes from the remote and checks the mergeability of the branches in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(base_checkCmd).Standalone()

	base_checkCmd.Flags().BoolP("help", "h", false, "Print help")
	baseCmd.AddCommand(base_checkCmd)
}
