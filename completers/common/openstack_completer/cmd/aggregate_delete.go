package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete existing aggregate(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_deleteCmd).Standalone()

	aggregateCmd.AddCommand(aggregate_deleteCmd)
}
