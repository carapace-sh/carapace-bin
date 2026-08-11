package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:     "sync",
	Short:   "Synchronize to a repository state",
	Aliases: []string{"synchronize"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().String("dependency-depth-limit", "0", "Maximum dependency traversal depth (0 means unlimited)")
	syncCmd.Flags().Bool("dependency-recursive", false, "Follow transitive dependencies recursively during dependency-based sync")
	syncCmd.Flags().StringSlice("dependency-tag", nil, "Tags to filter dependencies by during dependency-based sync")
	syncCmd.Flags().Bool("forward-changes", false, "Fast forward any local changes if syncing to a local revision")
	syncCmd.Flags().BoolP("help", "h", false, "Print help")
	syncCmd.Flags().Bool("reset", false, "Reset any local modified files to match the incoming revision")
	syncCmd.Flags().StringSlice("root-file", nil, "Root files for dependency-based selective sync (only sync changes for these files and their dependencies)")
	rootCmd.AddCommand(syncCmd)

	carapace.Gen(syncCmd).FlagCompletion(carapace.ActionMap{
		"root-file": carapace.ActionFiles(),
	})

	carapace.Gen(syncCmd).PositionalCompletion(
		action.ActionRevisions(syncCmd),
	)
}
