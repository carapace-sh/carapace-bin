package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run JavaScript with bun, a package.json script, or a bin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
