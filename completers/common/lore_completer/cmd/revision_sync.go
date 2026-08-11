package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_syncCmd = &cobra.Command{
	Use:     "sync",
	Short:   "Synchronize to a given state of a repository",
	Aliases: []string{"synchronize"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_syncCmd).Standalone()

	revision_syncCmd.Flags().String("dependency-depth-limit", "0", "Maximum dependency traversal depth (0 means unlimited)")
	revision_syncCmd.Flags().Bool("dependency-recursive", false, "Follow transitive dependencies recursively during dependency-based sync")
	revision_syncCmd.Flags().StringSlice("dependency-tag", nil, "Tags to filter dependencies by during dependency-based sync")
	revision_syncCmd.Flags().Bool("forward-changes", false, "Fast forward any local changes if syncing to a local revision")
	revision_syncCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_syncCmd.Flags().Bool("reset", false, "Reset any local modified files to match the incoming revision")
	revision_syncCmd.Flags().StringSlice("root-file", nil, "Root files for dependency-based selective sync (only sync changes for these files and their dependencies)")
	revisionCmd.AddCommand(revision_syncCmd)
}
