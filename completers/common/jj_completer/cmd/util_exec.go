package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var util_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute an external command via jj",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_execCmd).Standalone()

	util_execCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_execCmd)

	carapace.Gen(util_execCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(util_execCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
