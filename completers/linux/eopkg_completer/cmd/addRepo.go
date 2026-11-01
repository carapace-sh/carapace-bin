package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addRepoCmd = &cobra.Command{
	Use:     "add-repo <repo-name> <repo-uri>",
	Aliases: []string{"ar"},
	Short:   "add a new repository to the system with the given name and URI",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addRepoCmd).Standalone()

	addRepoCmd.Flags().Int("at", 0, "insert the new repository at the given index position")
	addRepoCmd.Flags().Bool("ignore-check", false, "ignore checking metadata for a valid distribution specifier")
	addRepoCmd.Flags().Bool("no-fetch", false, "do not download index, just register the new repository")

	rootCmd.AddCommand(addRepoCmd)
}
