package cmd

import (
	"github.com/spf13/cobra"
)

var workflow_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workflows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	workflowCmd.AddCommand(workflow_listCmd)

	workflowCmd.Flags().IntP("limit", "L", 50, "Maximum number of workflows to fetch")
	workflowCmd.Flags().BoolP("all", "a", false, "Show all workflows, including disabled workflows")
}
