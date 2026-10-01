package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Print the script to activate mise in an interactive shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(activateCmd).Standalone()

	activateCmd.Flags().BoolP("quiet", "q", false, "Suppress non-error messages")
	activateCmd.Flags().Bool("shims", false, "Use shims instead of PATH")
	activateCmd.Flags().Bool("status", false, "Show status of active mise environment")
	rootCmd.AddCommand(activateCmd)

	carapace.Gen(activateCmd).PositionalCompletion(
		action.ActionShells(),
	)
}
