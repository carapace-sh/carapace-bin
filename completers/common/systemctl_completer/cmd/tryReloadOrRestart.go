package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tryReloadOrRestartCmd = &cobra.Command{
	Use:     "try-reload-or-restart",
	Short:   "If active, reload one or more units, if supported, otherwise restart",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tryReloadOrRestartCmd).Standalone()

	rootCmd.AddCommand(tryReloadOrRestartCmd)

	carapace.Gen(tryReloadOrRestartCmd).PositionalAnyCompletion(
		action.ActionUnits(tryReloadOrRestartCmd).FilterArgs(),
	)
}
