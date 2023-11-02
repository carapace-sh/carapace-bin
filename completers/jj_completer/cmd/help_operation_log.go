package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_operation_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show the operation log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_operation_logCmd).Standalone()

	help_operationCmd.AddCommand(help_operation_logCmd)
}
