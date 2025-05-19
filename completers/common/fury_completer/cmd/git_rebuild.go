package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_rebuildCmd = &cobra.Command{
	Use:   "rebuild REPO",
	Short: "Run the builder on the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_rebuildCmd).Standalone()

	git_rebuildCmd.Flags().StringP("revision", "r", "", "Revision")
	gitCmd.AddCommand(git_rebuildCmd)
}
