package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_getCmd = &cobra.Command{
	Use:   "get <controller-id> [flags]",
	Short: "Get details of a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_getCmd).Standalone()

	runnerController_getCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_getCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerControllerCmd.AddCommand(runnerController_getCmd)
}
