package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_operationCmd = &cobra.Command{
	Use:   "operation",
	Short: "Commands for working with the operation log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_operationCmd).Standalone()

	helpCmd.AddCommand(help_operationCmd)
}
