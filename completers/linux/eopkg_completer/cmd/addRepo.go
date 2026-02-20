package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addRepoCmd = &cobra.Command{
	Use:     "add-repo",
	Short:   "Add a repository",
	Aliases: []string{"ar"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addRepoCmd).Standalone()

	addRepoCmd.Flags().String("at", "", "Add repository at given position")
	addRepoCmd.Flags().Bool("ignore-check", false, "Ignore repository distribution check")

	rootCmd.AddCommand(addRepoCmd)
}
