package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var isActiveCmd = &cobra.Command{
	Use:     "is-active",
	Short:   "Check whether units are active",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isActiveCmd).Standalone()

	rootCmd.AddCommand(isActiveCmd)

	carapace.Gen(isActiveCmd).PositionalAnyCompletion(
		action.ActionUnits(isActiveCmd).FilterArgs(),
	)
}
