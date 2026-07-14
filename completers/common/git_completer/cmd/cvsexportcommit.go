package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var cvsexportcommitCmd = &cobra.Command{
	Use:     "cvsexportcommit",
	Short:   "Export a single commit to a CVS checkout",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(cvsexportcommitCmd).Standalone()

	cvsexportcommitCmd.Flags().BoolS("P", "P", false, "force the parent commit, even if it is not a direct parent")
	cvsexportcommitCmd.Flags().BoolS("W", "W", false, "tell cvsexportcommit that the current working directory is not only a Git checkout, but also the CVS checkout")
	cvsexportcommitCmd.Flags().BoolS("a", "a", false, "add authorship information")
	cvsexportcommitCmd.Flags().BoolS("c", "c", false, "commit automatically if the patch applied cleanly")
	cvsexportcommitCmd.Flags().BoolS("d", "d", false, "set an alternative CVSROOT to use")
	cvsexportcommitCmd.Flags().BoolS("f", "f", false, "force the merge even if the files are not up to date")
	cvsexportcommitCmd.Flags().BoolS("k", "k", false, "reverse CVS keyword expansion in working CVS checkout before applying patch")
	cvsexportcommitCmd.Flags().BoolS("m", "m", false, "prepend the commit message with the provided prefix")
	cvsexportcommitCmd.Flags().BoolS("p", "p", false, "be pedantic (paranoid) when applying patches")
	cvsexportcommitCmd.Flags().BoolS("u", "u", false, "update affected files from CVS repository before attempting export")
	cvsexportcommitCmd.Flags().BoolS("v", "v", false, "verbose")
	cvsexportcommitCmd.Flags().BoolS("w", "w", false, "specify the location of the CVS checkout to use for the export")
	rootCmd.AddCommand(cvsexportcommitCmd)

	carapace.Gen(cvsexportcommitCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()), // TODO verify
		git.ActionRefs(git.RefOption{}.Default()), // TODO verify
	)
}
