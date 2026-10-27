package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:     "branch",
	Short:   "List branches, or create / delete / rename / configure one",
	GroupID: "branches",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().StringSliceS("M", "M", nil, "Force rename: `-M <new>` or `-M <old> <new>`")
	branchCmd.Flags().BoolP("all", "a", false, "List local and remote-tracking branches")
	branchCmd.Flags().String("contains", "", "List branches whose tip contains `REV` (defaults to `HEAD` when omitted)")
	branchCmd.Flags().BoolP("delete", "d", false, "Delete the named branch instead of creating it")
	branchCmd.Flags().Bool("dry-run", false, "Preview the operation without updating refs, config, or reflogs")
	branchCmd.Flags().BoolP("force", "f", false, "Overwrite an existing branch when creating or renaming")
	branchCmd.Flags().BoolP("force-delete", "D", false, "Delete even when the branch is not fully merged into the current branch")
	branchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.Flags().StringP("limit", "n", "", "Maximum number of rows to list")
	branchCmd.Flags().StringSlice("list", nil, "Include only branches whose short name matches `GLOB` (repeatable)")
	branchCmd.Flags().String("max-count", "", "Maximum number of rows to list")
	branchCmd.Flags().String("merged", "", "List branches merged into `REV` (defaults to `HEAD` when omitted)")
	branchCmd.Flags().StringSliceP("move", "m", nil, "Rename a branch: `-m <new>` or `-m <old> <new>`")
	branchCmd.Flags().String("no-merged", "", "List branches not merged into `REV` (defaults to `HEAD` when omitted)")
	branchCmd.Flags().Bool("no-track", false, "Do not configure upstream when creating")
	branchCmd.Flags().BoolP("remotes", "r", false, "List remote-tracking branches under `refs/remotes/`")
	branchCmd.Flags().Bool("reverse", false, "Reverse sort order")
	branchCmd.Flags().StringP("set-upstream-to", "u", "", "Set upstream for the named branch (or the current branch when omitted)")
	branchCmd.Flags().Bool("show-current", false, "Print the current branch name only (same as `git branch --show-current`)")
	branchCmd.Flags().String("sort", "", "Sort key: `name`, `version`, or `date`")
	branchCmd.Flags().BoolP("track", "t", false, "Configure upstream from the start point when creating")
	branchCmd.Flags().Bool("unset-upstream", false, "Remove upstream configuration for the named branch (or the current branch)")
	rootCmd.AddCommand(branchCmd)

	carapace.Gen(branchCmd).FlagCompletion(carapace.ActionMap{
		"contains":        git.ActionRefs(git.RefOption{}.Default()),
		"merged":          git.ActionRefs(git.RefOption{}.Default()),
		"no-merged":       git.ActionRefs(git.RefOption{}.Default()),
		"set-upstream-to": git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}),
		"sort":            carapace.ActionValues("date", "name", "version"),
	})

	carapace.Gen(branchCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{LocalBranches: true}),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
