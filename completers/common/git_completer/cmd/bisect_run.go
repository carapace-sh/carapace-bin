package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var bisect_runCmd = &cobra.Command{
	Use:   "run",
	Short: "run script to automatically bisect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_runCmd).Standalone()
	bisect_runCmd.Flags().SetInterspersed(false)

	bisectCmd.AddCommand(bisect_runCmd)

	carapace.Gen(bisect_runCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionExecutables(),
		).ToA(),
	)

	carapace.Gen(bisect_runCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	carapace.Gen(bisect_runCmd).DashAnyCompletion(
		carapace.ActionPositional(bisect_runCmd),
	)
}
