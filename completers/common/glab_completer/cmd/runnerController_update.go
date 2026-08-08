package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_updateCmd = &cobra.Command{
	Use:   "update <id> [flags]",
	Short: "Update a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_updateCmd).Standalone()

	runnerController_updateCmd.Flags().StringP("description", "d", "", "Description of the runner controller.")
	runnerController_updateCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_updateCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerController_updateCmd.Flags().String("state", "", "State of the runner controller: disabled, enabled, dry_run.")
	runnerControllerCmd.AddCommand(runnerController_updateCmd)

	carapace.Gen(runnerController_updateCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
		"state":  carapace.ActionValues("disabled", "enabled", "dry_run"),
	})
}
