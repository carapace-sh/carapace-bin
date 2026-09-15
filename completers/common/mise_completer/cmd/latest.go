package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Gets the latest available version for a tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(latestCmd).Standalone()

	rootCmd.AddCommand(latestCmd)

	carapace.Gen(latestCmd).PositionalCompletion(
		action.ActionTools(),
	)
}
