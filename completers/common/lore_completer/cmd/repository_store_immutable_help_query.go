package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_store_immutable_help_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_store_immutable_help_queryCmd).Standalone()

	repository_store_immutable_helpCmd.AddCommand(repository_store_immutable_help_queryCmd)
}
