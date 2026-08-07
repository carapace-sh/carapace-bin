package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List runner controllers. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_listCmd).Standalone()

	runnerController_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerController_listCmd.Flags().StringP("page", "p", "", "Page number.")
	runnerController_listCmd.Flags().StringP("per-page", "P", "", "Number of items per page.")
	runnerControllerCmd.AddCommand(runnerController_listCmd)
}
