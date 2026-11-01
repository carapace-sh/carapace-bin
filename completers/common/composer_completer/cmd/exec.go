package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executes a vendored binary/script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().BoolP("list", "l", false, "List the available binaries")

	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).PositionalCompletion(
		action.ActionBinaries(execCmd),
	)

	carapace.Gen(execCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
