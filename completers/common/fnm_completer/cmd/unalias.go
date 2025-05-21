package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unaliasCmd = &cobra.Command{
	Use:   "unalias",
	Short: "Remove an alias definition",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(unaliasCmd).Standalone()

	rootCmd.AddCommand(unaliasCmd)

	carapace.Gen(unaliasCmd).PositionalAnyCompletion(
		action.ActionAliases(),
	)
}
