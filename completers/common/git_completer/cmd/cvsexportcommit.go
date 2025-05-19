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

	cvsexportcommitCmd.Flags().BoolS("P", "P", false, "Force the parent commit, even if it is not a direct parent")
	cvsexportcommitCmd.Flags().BoolS("W", "W", false, "Tell cvsexportcommit that the current working directory is not only a Git checkout, but also the CVS checkout")
	cvsexportcommitCmd.Flags().BoolS("a", "a", false, "Add authorship information")
	cvsexportcommitCmd.Flags().BoolS("c", "c", false, "Commit automatically if the patch applied cleanly")
	cvsexportcommitCmd.Flags().BoolS("d", "d", false, "Set an alternative CVSROOT to use")
	cvsexportcommitCmd.Flags().BoolS("f", "f", false, "Force the merge even if the files are not up to date")
	cvsexportcommitCmd.Flags().BoolS("k", "k", false, "Reverse CVS keyword expansion in working CVS checkout before applying patch")
	cvsexportcommitCmd.Flags().BoolS("m", "m", false, "Prepend the commit message with the provided prefix")
	cvsexportcommitCmd.Flags().BoolS("p", "p", false, "Be pedantic (paranoid) when applying patches")
	cvsexportcommitCmd.Flags().BoolS("u", "u", false, "Update affected files from CVS repository before attempting export")
	cvsexportcommitCmd.Flags().BoolS("v", "v", false, "Verbose")
	cvsexportcommitCmd.Flags().BoolS("w", "w", false, "Specify the location of the CVS checkout to use for the export")
	rootCmd.AddCommand(cvsexportcommitCmd)

	carapace.Gen(cvsexportcommitCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()), // TODO verify
		git.ActionRefs(git.RefOption{}.Default()), // TODO verify
	)
}
