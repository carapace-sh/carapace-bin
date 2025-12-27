package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show properties of one or more units/jobs or the manager",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalAnyCompletion(
		action.ActionUnits(showCmd).FilterArgs(),
	)
}
