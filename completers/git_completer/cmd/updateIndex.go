package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var updateIndexCmd = &cobra.Command{
	Use:     "update-index",
	Short:   "Register file contents in the working tree to the index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(updateIndexCmd).Standalone()

	updateIndexCmd.Flags().Bool("add", false, "if a specified file isn’t in the index already then it’s added")
	updateIndexCmd.Flags().BoolP("again", "g", false, "runs git update-index itself on the paths whose index entries differ from HEAD")
	updateIndexCmd.Flags().Bool("assume-unchanged", false, "set the \"assume unchanged\" bit for paths")
	updateIndexCmd.Flags().String("cacheinfo", "", "directly insert the specified info into the index")
	updateIndexCmd.Flags().String("chmod", "", "set the execute permissions on the updated files")
	updateIndexCmd.Flags().Bool("force-remove", false, "remove the file from the index even when the working directory still has such a file")
	updateIndexCmd.Flags().Bool("force-untracked-cache", false, "same as --untracked-cache")
	updateIndexCmd.Flags().Bool("fsmonitor", false, "enable files system monitor feature")
	updateIndexCmd.Flags().Bool("fsmonitor-valid", false, "set the \"fsmonitor valid\" bit for paths")
	updateIndexCmd.Flags().Bool("ignore-missing", false, "ignores missing files during a --refresh")
	updateIndexCmd.Flags().Bool("ignore-skip-worktree-entries", false, "do not remove skip-worktree (AKA \"index-only\") entries")
	updateIndexCmd.Flags().Bool("ignore-submodules", false, "do not try to update submodules")
	updateIndexCmd.Flags().Bool("index-info", false, "read index information from stdin")
	updateIndexCmd.Flags().String("index-version", "", "write the resulting index out in the named on-disk format version")
	updateIndexCmd.Flags().Bool("info-only", false, "just insert object IDs into the index for following files")
	updateIndexCmd.Flags().Bool("no-assume-unchanged", false, "unset the \"assume unchanged\" bit for paths")
	updateIndexCmd.Flags().Bool("no-fsmonitor", false, "disable files system monitor feature")
	updateIndexCmd.Flags().Bool("no-fsmonitor-valid", false, "unset the \"fsmonitor valid\" bit for paths")
	updateIndexCmd.Flags().Bool("no-ignore-skip-worktree-entries", false, "remove skip-worktree (AKA \"index-only\") entries")
	updateIndexCmd.Flags().Bool("no-skip-worktree", false, "unset the \"skip-worktree\" bit for paths")
	updateIndexCmd.Flags().Bool("no-split-index", false, "disable split index mode")
	updateIndexCmd.Flags().Bool("no-untracked-cache", false, "disable untracked cache feature")
	updateIndexCmd.Flags().BoolS("q", "q", false, "quiet")
	updateIndexCmd.Flags().Bool("really-refresh", false, "like --refresh, but checks stat information unconditionally")
	updateIndexCmd.Flags().Bool("refresh", false, "looks at the current index and checks to see if merges or updates are needed")
	updateIndexCmd.Flags().Bool("remove", false, "if a specified file is in the index but is missing then it’s removed")
	updateIndexCmd.Flags().Bool("replace", false, "existing entries that conflict with the entry being added are automatically removed")
	updateIndexCmd.Flags().Bool("show-index-version", false, "report the index format version used by the on-disk index file")
	updateIndexCmd.Flags().Bool("skip-worktree", false, "set the \"skip-worktree\" bit for paths")
	updateIndexCmd.Flags().Bool("split-index", false, "enable split index mode")
	updateIndexCmd.Flags().Bool("stdin", false, "read a list of paths from the standard input")
	updateIndexCmd.Flags().Bool("test-untracked-cache", false, "only perform tests on the working directory to make sure untracked cache can be used")
	updateIndexCmd.Flags().Bool("unmerged", false, "continue if --refresh finds unmerged changes in the index")
	updateIndexCmd.Flags().Bool("unresolve", false, "restores the unmerged or needs updating state of a file during a merge")
	updateIndexCmd.Flags().Bool("untracked-cache", false, "enable untracked cache feature")
	updateIndexCmd.Flags().Bool("verbose", false, "report what is being added and removed from the index")
	updateIndexCmd.Flags().BoolS("z", "z", false, "paths are separated with NUL character instead of LF")
	rootCmd.AddCommand(updateIndexCmd)

	carapace.Gen(updateIndexCmd).FlagCompletion(carapace.ActionMap{
		"chmod": carapace.ActionStyledValues(
			"+x", style.Carapace.KeywordPositive,
			"-x", style.Carapace.KeywordNegative,
		),
		"index-version": carapace.ActionValues("2", "3", "4"),
	})

	carapace.Gen(updateIndexCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(updateIndexCmd).DashAnyCompletion(
		carapace.ActionPositional(updateIndexCmd),
	)
}
