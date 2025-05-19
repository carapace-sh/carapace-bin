package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var diffIndexCmd = &cobra.Command{
	Use:     "diff-index",
	Short:   "Compare a tree to the working tree or index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(diffIndexCmd).Standalone()

	diffIndexCmd.Flags().StringS("G", "G", "", "look for differences whose patch text contains added/removed lines that match <regex>")
	diffIndexCmd.Flags().Bool("default-prefix", false, "use the default source and destination prefixes")
	diffIndexCmd.Flags().StringP("ignore-matching-lines", "I", "", "ignore changes whose all lines match <regex>")
	diffIndexCmd.Flags().BoolS("m", "m", false, "make git diff-index say that all non-checked-out files are up to date")
	diffIndexCmd.Flags().Bool("merge-base", false, "use the merge base between <tree-ish> and HEAD")
	diffIndexCmd.Flags().Bool("no-relative", false, "do not make the output relative")
	diffIndexCmd.Flags().Bool("no-textconv", false, "do not external text conversion filters to be run")
	diffIndexCmd.Flags().Bool("quiet", false, "disable all output of the program")
	diffIndexCmd.Flags().String("rotate-to", "", "move files before the named <file> to the end")
	diffIndexCmd.Flags().String("skip-to", "", "discard the files before the named <file> from the output ")
	common.AddDiffFlags(diffIndexCmd)
	rootCmd.AddCommand(diffIndexCmd)

	carapace.Gen(diffIndexCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		carapace.ActionFiles(),
	)
}
