package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_logsCmd).Standalone()
	apps_logsCmd.Flags().String("deployment", "", "The deployment ID. Defaults to current deployment.")
	apps_logsCmd.Flags().BoolP("follow", "f", false, "Follow logs as they are emitted.")
	apps_logsCmd.Flags().Int("tail", -1, "Number of lines to show from the end of the log.")
	apps_logsCmd.Flags().String("type", "run", "The type of logs.")
	appsCmd.AddCommand(apps_logsCmd)

	// TODO deployment completion
	carapace.Gen(apps_logsCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("build", "deploy", "run"),
	})

	// TODO positional
}
