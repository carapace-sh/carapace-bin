package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var maskCmd = &cobra.Command{
	Use:     "mask",
	Short:   "Mask one or more units",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maskCmd).Standalone()

	rootCmd.AddCommand(maskCmd)

	carapace.Gen(maskCmd).PositionalAnyCompletion(
		action.ActionUnits(maskCmd).FilterArgs(),
	)
}
