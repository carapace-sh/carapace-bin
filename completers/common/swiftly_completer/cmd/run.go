package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <command>...",
	Short: "Run a command while proxying to the selected toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("help", "h", false, "Show help information")

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
