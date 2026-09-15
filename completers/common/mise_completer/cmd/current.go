package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Shows current active and installed versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(currentCmd).Standalone()

	rootCmd.AddCommand(currentCmd)

	carapace.Gen(currentCmd).PositionalCompletion(
		action.ActionTools(),
	)
}
