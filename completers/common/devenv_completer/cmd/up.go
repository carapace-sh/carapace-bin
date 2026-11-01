package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start processes in the foreground",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().BoolP("detach", "d", false, "Start processes in the background")
	upCmd.Flags().StringP("mode", "m", "", "The execution mode for process tasks (affects dependency resolution)")
	upCmd.Flags().Bool("no-strict-ports", false, "Disable strict port mode, overriding strict_ports from devenv.yaml")
	upCmd.Flags().Bool("strict-ports", false, "Error if a port is already in use instead of auto-allocating the next available port")

	rootCmd.AddCommand(upCmd)

	upCmd.MarkFlagsMutuallyExclusive("no-strict-ports", "strict-ports")

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"mode": actionExecutionModes(),
	})

	carapace.Gen(upCmd).PositionalAnyCompletion(
		action.ActionProcesses(upCmd).FilterArgs(),
	)
}
