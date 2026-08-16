package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Print shell hook for auto-activation on directory change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookCmd).Standalone()

	rootCmd.AddCommand(hookCmd)

	carapace.Gen(hookCmd).PositionalCompletion(
		actionShells(),
	)
}
