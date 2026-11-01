package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runner_updateCmd = &cobra.Command{
	Use:   "update <runner-id>",
	Short: "Update a runner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_updateCmd).Standalone()

	runner_updateCmd.Flags().Bool("pause", false, "Pause the runner.")
	runner_updateCmd.Flags().Bool("unpause", false, "Resume a paused runner.")
	runnerCmd.AddCommand(runner_updateCmd)
}
