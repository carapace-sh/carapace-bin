package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_cacheCmd).Standalone()

	aggregateCmd.AddCommand(aggregate_cacheCmd)
}
