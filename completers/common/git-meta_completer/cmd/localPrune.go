package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var localPruneCmd = &cobra.Command{
	Use:    "local-prune",
	Short:  "Prune old metadata from the local SQLite database",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localPruneCmd).Standalone()

	localPruneCmd.Flags().Bool("dry-run", false, "Show what would be pruned without deleting anything")
	localPruneCmd.Flags().BoolP("help", "h", false, "Print help")
	localPruneCmd.Flags().Bool("skip-date", false, "Ignore the date rule and prune all non-project metadata")
	rootCmd.AddCommand(localPruneCmd)
}
