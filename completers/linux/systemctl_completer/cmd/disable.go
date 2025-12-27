package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:     "disable",
	Short:   "Disable one or more unit files",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableCmd).Standalone()

	rootCmd.AddCommand(disableCmd)

	carapace.Gen(disableCmd).PositionalAnyCompletion(
		action.ActionUnits(disableCmd).FilterArgs(),
	)
}
