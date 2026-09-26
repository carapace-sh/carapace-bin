package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:    "clean",
	Short:  "Clear the cache, removing all entries or those linked to specific packages",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("force", false, "Force removal of the cache, ignoring in-use checks")
	rootCmd.AddCommand(cleanCmd)
}
