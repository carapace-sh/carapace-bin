package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workItemsCmd = &cobra.Command{
	Use:   "work-items <command> [flags]",
	Short: "Manage work items. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workItemsCmd).Standalone()

	rootCmd.AddCommand(workItemsCmd)
}
