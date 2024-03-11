package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terramate"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("cloud-status", "", "Filter by status")
	listCmd.Flags().String("experimental-status", "", "Filter by status (Deprecated)")
	listCmd.Flags().Bool("run-order", false, "Sort stacks by order of execution")
	listCmd.Flags().Bool("why", false, "Shows the reason why the stack has changed")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"cloud-status": terramate.ActionCloudStatus(),
	})
}
