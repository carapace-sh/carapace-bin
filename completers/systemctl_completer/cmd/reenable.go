package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reenableCmd = &cobra.Command{
	Use:     "reenable",
	Short:   "Reenable one or more unit files",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reenableCmd).Standalone()

	rootCmd.AddCommand(reenableCmd)

	carapace.Gen(reenableCmd).PositionalAnyCompletion(
		action.ActionUnits(reenableCmd).FilterArgs(),
	)
}
