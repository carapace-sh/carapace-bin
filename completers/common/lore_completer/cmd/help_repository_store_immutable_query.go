package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_store_immutable_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_store_immutable_queryCmd).Standalone()

	help_repository_store_immutableCmd.AddCommand(help_repository_store_immutable_queryCmd)
}
