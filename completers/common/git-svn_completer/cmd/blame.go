package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Show what revision and author last modified each line of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()

	blameCmd.Flags().Bool("git-format", false, "Produce output in git blame format with SVN revision numbers")
	rootCmd.AddCommand(blameCmd)

	carapace.Gen(blameCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
