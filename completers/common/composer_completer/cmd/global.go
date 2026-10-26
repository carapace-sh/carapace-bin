package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Allows running commands in the global composer dir ($COMPOSER_HOME)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalCmd).Standalone()
	globalCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(globalCmd)

	carapace.Gen(globalCmd).PositionalCompletion(
		action.ActionCommands(),
	)

	carapace.Gen(globalCmd).PositionalAnyCompletion(
		action.ActionGlobal(),
	)
}
