package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var patchCommitCmd = &cobra.Command{
	Use:     "patch-commit",
	Short:   "Generate a patch out of a directory",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCommitCmd).Standalone()

	patchCommitCmd.Flags().BoolP("save", "s", false, "Add the patch to your resolution entries")
	rootCmd.AddCommand(patchCommitCmd)

	carapace.Gen(patchCommitCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
