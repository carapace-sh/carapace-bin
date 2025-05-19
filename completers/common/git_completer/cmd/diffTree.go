package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var diffTreeCmd = &cobra.Command{
	Use:     "diff-tree",
	Short:   "diff-tree Compares the content and mode of blobs found via two tree objects",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(diffTreeCmd).Standalone()

	diffTreeCmd.Flags().StringS("G", "G", "", "look for differences whose patch text contains added/removed lines that match <regex>")
	diffTreeCmd.Flags().Bool("always", false, "show the commit itself and the commit log message even if the diff itself is empty")
	diffTreeCmd.Flags().Bool("combined-all-paths", false, "let combined diffs list the name of the file from all parents")
	diffTreeCmd.Flags().Bool("default-prefix", false, "use the default source and destination prefixes")
	diffTreeCmd.Flags().StringP("ignore-matching-lines", "I", "", "ignore changes whose all lines match <regex>")
	diffTreeCmd.Flags().BoolS("m", "m", false, "show differences for merge commits")
	diffTreeCmd.Flags().Bool("merge-base", false, "use the merge base between the two <tree-ish>s as the \"before\" side")
	diffTreeCmd.Flags().Bool("no-commit-id", false, "suppress commit ID output")
	diffTreeCmd.Flags().Bool("no-notes", false, "do not show notes")
	diffTreeCmd.Flags().Bool("no-relative", false, "do not show relative pathnames")
	diffTreeCmd.Flags().Bool("no-textconv", false, "do not allow external text conversion filters")
	diffTreeCmd.Flags().String("notes", "", "show the notes that annotate the commit")
	diffTreeCmd.Flags().Bool("quiet", false, "disable all output of the program")
	diffTreeCmd.Flags().BoolS("r", "r", false, "recurse into sub-trees")
	diffTreeCmd.Flags().Bool("root", false, "show the initial commit as a big creation event")
	diffTreeCmd.Flags().String("rotate-to", "", "move the files before the named <file> to end")
	diffTreeCmd.Flags().Bool("show-notes-by-default", false, "show the default notes unless options for displaying specific notes are given")
	diffTreeCmd.Flags().String("skip-to", "", "discard the files before the named <file> from the output")
	diffTreeCmd.Flags().Bool("stdin", false, "read from standard input")
	diffTreeCmd.Flags().BoolS("t", "t", false, "show tree entry itself as well as subtrees")
	diffTreeCmd.Flags().BoolS("v", "v", false, "also show the commit message before the differences")
	common.AddDiffFlags(diffTreeCmd)
	common.AddCommitFormattingOptions(diffTreeCmd)
	rootCmd.AddCommand(diffTreeCmd)

	diffTreeCmd.Flag("notes").NoOptDefVal = " "
	diffTreeCmd.Flag("relative").NoOptDefVal = " "

	carapace.Gen(diffTreeCmd).FlagCompletion(carapace.ActionMap{
		"notes":    git.ActionRefs(git.RefOption{}.Default()),
		"relative": carapace.ActionDirectories(),
	})

	carapace.Gen(diffTreeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(diffTreeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
