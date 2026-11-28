package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

var setTargetCmd = &cobra.Command{
	Use:   "set-target",
	Short: "Set the chip target to build",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setTargetCmd).Standalone()
	rootCmd.AddCommand(setTargetCmd)
	carapace.Gen(setTargetCmd).PositionalCompletion(
		espidf.ActionTargets(),
	)
}
