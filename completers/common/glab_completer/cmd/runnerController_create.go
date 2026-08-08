package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_createCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_createCmd).Standalone()

	runnerController_createCmd.Flags().StringP("description", "d", "", "Description of the runner controller.")
	runnerController_createCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_createCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerController_createCmd.Flags().String("state", "", "State of the runner controller: disabled, enabled, dry_run.")
	runnerControllerCmd.AddCommand(runnerController_createCmd)

	carapace.Gen(runnerController_createCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
		"state":  carapace.ActionValues("disabled", "enabled", "dry_run"),
	})
}
