package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var isEnabledCmd = &cobra.Command{
	Use:     "is-enabled",
	Short:   "Check whether unit files are enabled",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isEnabledCmd).Standalone()

	rootCmd.AddCommand(isEnabledCmd)

	carapace.Gen(isEnabledCmd).PositionalAnyCompletion(
		action.ActionUnits(isEnabledCmd).FilterArgs(),
	)
}
