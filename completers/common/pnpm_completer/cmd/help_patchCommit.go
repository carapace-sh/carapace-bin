package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_patchCommitCmd = &cobra.Command{
	Use:   "patch-commit",
	Short: "Generate a patch out of a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_patchCommitCmd).Standalone()

	helpCmd.AddCommand(help_patchCommitCmd)
}
