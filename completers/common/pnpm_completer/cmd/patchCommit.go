package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var patchCommitCmd = &cobra.Command{
	Use:   "patch-commit",
	Short: "Generate a patch out of a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCommitCmd).Standalone()

	patchCommitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	patchCommitCmd.Flags().String("patches-dir", "", "The generated patch file will be saved to this directory")

	carapace.Gen(patchCommitCmd).FlagCompletion(carapace.ActionMap{
		"patches-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(patchCommitCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	rootCmd.AddCommand(patchCommitCmd)
}
