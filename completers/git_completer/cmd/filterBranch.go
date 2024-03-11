package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var filterBranchCmd = &cobra.Command{
	Use:     "filter-branch",
	Short:   "Rewrite branches",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {
	carapace.Gen(filterBranchCmd).Standalone()

	filterBranchCmd.Flags().String("commit-filter", "", "This is the filter for performing the commit")
	filterBranchCmd.Flags().StringS("d", "d", "", "Use this option to set the path to the temporary directory used for rewriting")
	filterBranchCmd.Flags().String("env-filter", "", "This filter may be used if you only need to modify the environment in which the commit will be performed")
	filterBranchCmd.Flags().BoolP("force", "f", false, "Force to start with an existing temporary directory or when there are already refs starting with refs/original/")
	filterBranchCmd.Flags().String("index-filter", "", "This is the filter for rewriting the index")
	filterBranchCmd.Flags().String("msg-filter", "", "This is the filter for rewriting the commit messages")
	filterBranchCmd.Flags().String("original", "", "Use this option to set the namespace where the original commits will be stored")
	filterBranchCmd.Flags().String("parent-filter", "", "This is the filter for rewriting the commit’s parent list")
	filterBranchCmd.Flags().Bool("prune-empty", false, "Some filters will generate empty commits that leave the tree untouched")
	filterBranchCmd.Flags().String("setup", "", "This is not a real filter executed for each commit but a one time setup just before the loop")
	filterBranchCmd.Flags().String("state-branch", "", "This option will cause the mapping from old to new objects to be loaded from named branch")
	filterBranchCmd.Flags().String("subdirectory-filter", "", "Only look at the history which touches the given subdirectory")
	filterBranchCmd.Flags().String("tag-name-filter", "", "This is the filter for rewriting tag names")
	filterBranchCmd.Flags().String("tree-filter", "", "This is the filter for rewriting the tree and its contents")
	rootCmd.AddCommand(filterBranchCmd)

	carapace.Gen(filterBranchCmd).FlagCompletion(carapace.ActionMap{
		"commit-filter":       bridge.ActionCarapaceBin().SplitP(),
		"d":                   carapace.ActionDirectories(),
		"env-filter":          bridge.ActionCarapaceBin().SplitP(),
		"index-filter":        bridge.ActionCarapaceBin().SplitP(),
		"msg-filter":          bridge.ActionCarapaceBin().SplitP(),
		"parent-filter":       bridge.ActionCarapaceBin().SplitP(),
		"setup":               bridge.ActionCarapaceBin().SplitP(),
		"state-branch":        git.ActionLocalBranches(),
		"subdirectory-filter": carapace.ActionDirectories(),
		"tag-name-filter":     bridge.ActionCarapaceBin().SplitP(),
		"tree-filter":         bridge.ActionCarapaceBin().SplitP(),
	})

	carapace.Gen(filterBranchCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin("git", "rev-list"),
	)
}
