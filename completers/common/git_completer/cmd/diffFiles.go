package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var diffFilesCmd = &cobra.Command{
	Use:     "diff-files",
	Short:   "Compares files in the working tree and the index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(diffFilesCmd).Standalone()

	diffFilesCmd.Flags().StringS("G", "G", "", "Look for differences whose patch text contains added/removed lines that match <regex>")
	diffFilesCmd.Flags().Bool("default", false, "Use argument as default revision")
	diffFilesCmd.Flags().Bool("default-prefix", false, "Use the default source and destination prefixes")
	diffFilesCmd.Flags().Bool("early-output", false, "undocumented")
	diffFilesCmd.Flags().Bool("follow", false, "Continue listing the history of a file beyond renames")
	diffFilesCmd.Flags().Bool("full-diff", false, "Show full commit diffs")
	diffFilesCmd.Flags().StringP("ignore-matching-lines", "I", "", "Ignore changes whose all lines match <regex>")
	diffFilesCmd.Flags().Bool("log-size", false, "Print log message size in bytes before the message")
	diffFilesCmd.Flags().Bool("mailmap", false, "Use mailmap file to map author and committer name")
	diffFilesCmd.Flags().Bool("no-relative", false, "Do not show relative pathnames")
	diffFilesCmd.Flags().BoolS("q", "q", false, "Remain silent even for nonexistent files")
	diffFilesCmd.Flags().String("rotate-to", "", "Move the files before the named <file> to the end")
	diffFilesCmd.Flags().String("skip-to", "", "Discard the files before the named <file> from the output")
	diffFilesCmd.Flags().Bool("use-mailmap", false, "Use mailmap file to map author and committer name")
	common.AddDiffFlags(diffFilesCmd)
	common.AddPrettyFlags(diffFilesCmd)
	common.AddCommitLimitingOptions(diffFilesCmd)
	common.AddHistorySimplificationOptions(diffFilesCmd)
	common.AddBisectionHelperOptions(diffFilesCmd)
	common.AddCommitOrderingOptions(diffFilesCmd)
	common.AddObjectTraversalOptions(diffFilesCmd)
	common.AddCommitFormattingOptions(diffFilesCmd)
	rootCmd.AddCommand(diffFilesCmd)

	carapace.Gen(diffFilesCmd).PositionalAnyCompletion(
		git.ActionChanges(git.ChangeOpts{Unstaged: true}).FilterArgs(),
	)
}
