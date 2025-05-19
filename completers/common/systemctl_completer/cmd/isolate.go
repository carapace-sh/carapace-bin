package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var isolateCmd = &cobra.Command{
	Use:     "isolate",
	Short:   "Start one unit and stop all others",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isolateCmd).Standalone()

	rootCmd.AddCommand(isolateCmd)

	carapace.Gen(isolateCmd).PositionalCompletion(
		action.ActionUnits(isolateCmd),
	)
}
