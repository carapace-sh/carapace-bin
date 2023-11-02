package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operation_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show the operation log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_logCmd).Standalone()

	operation_logCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operation_logCmd.Flags().StringP("limit", "l", "", "Limit number of operations to show")
	operation_logCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of operations")
	operation_logCmd.Flags().StringP("template", "T", "", "Render each operation using the given template")
	operationCmd.AddCommand(operation_logCmd)
}
