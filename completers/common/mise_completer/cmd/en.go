package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var enCmd = &cobra.Command{
	Use:     "en",
	Short:   "Execute a command with tool(s) set",
	Aliases: []string{"exec", "x"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enCmd).Standalone()

	rootCmd.AddCommand(enCmd)

	carapace.Gen(enCmd).PositionalAnyCompletion(
		action.ActionTools(),
	)
}
