package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Start (activate) one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).PositionalAnyCompletion(
		action.ActionUnits(startCmd).FilterArgs(),
	)
}
