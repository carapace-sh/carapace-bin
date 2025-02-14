package cmd

import (
	"github.com/carapace-sh/carapace"
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
	operation_logCmd.Flags().StringP("limit", "n", "", "Limit number of revisions to show")
	operation_logCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of operations")
	operation_logCmd.Flags().Bool("op-diff", false, "Show changes to the repository at each operation")
	operation_logCmd.Flags().BoolP("patch", "p", false, "Show patch of modifications to changes (implies --op-diff)")
	operation_logCmd.Flags().Bool("reversed", false, " Show operations in the opposite order (older operations first)")
	operation_logCmd.Flags().StringP("template", "T", "", "Render each operation using the given template")
	operationCmd.AddCommand(operation_logCmd)
}
