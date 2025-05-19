package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:     "reload",
	Short:   "Reload one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadCmd).Standalone()

	rootCmd.AddCommand(reloadCmd)

	carapace.Gen(reloadCmd).PositionalAnyCompletion(
		action.ActionUnits(reloadCmd).FilterArgs(),
	)
}
