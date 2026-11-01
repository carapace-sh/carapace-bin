package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_store_help_immutable_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_store_help_immutable_queryCmd).Standalone()

	repository_store_help_immutableCmd.AddCommand(repository_store_help_immutable_queryCmd)
}
