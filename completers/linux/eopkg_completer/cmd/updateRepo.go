package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateRepoCmd = &cobra.Command{
	Use:     "update-repo",
	Short:   "Update repository databases",
	Aliases: []string{"ur"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateRepoCmd).Standalone()

	updateRepoCmd.Flags().BoolP("force", "f", false, "Force updating repos")

	rootCmd.AddCommand(updateRepoCmd)
}
