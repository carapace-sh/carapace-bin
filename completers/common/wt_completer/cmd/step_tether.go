package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var step_tetherCmd = &cobra.Command{
	Use:   "tether",
	Short: "[experimental] Run a command; kill its whole process tree when its worktree is removed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_tetherCmd).Standalone()

	step_tetherCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_tetherCmd)

	carapace.Gen(step_tetherCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
