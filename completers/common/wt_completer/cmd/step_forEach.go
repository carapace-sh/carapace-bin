package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var step_forEachCmd = &cobra.Command{
	Use:   "for-each",
	Short: "[experimental] Run command in each worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_forEachCmd).Standalone()

	step_forEachCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_forEachCmd)

	carapace.Gen(step_forEachCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
