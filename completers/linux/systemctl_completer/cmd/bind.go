package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bindCmd = &cobra.Command{
	Use:     "bind",
	Short:   "Bind-mount a path from the host into a unit's namespace",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bindCmd).Standalone()

	rootCmd.AddCommand(bindCmd)

	carapace.Gen(bindCmd).PositionalCompletion(
		action.ActionUnits(bindCmd),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
