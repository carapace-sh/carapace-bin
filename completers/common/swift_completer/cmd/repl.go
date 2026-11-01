package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Experiment with Swift code interactively",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replCmd).Standalone()
	replCmd.Flags().SetInterspersed(false)

	common.AddCompilerFlags(replCmd)
	common.CompilerFlagCompletion(replCmd)

	replCmd.Flags().BoolP("help", "h", false, "Show help information")
	replCmd.Flags().Bool("version", false, "Show the version")

	rootCmd.AddCommand(replCmd)

	carapace.Gen(replCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".swift"),
	)
}
