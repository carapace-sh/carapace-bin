package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_store_immutable_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_store_immutable_queryCmd).Standalone()

	repository_store_immutable_queryCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_store_immutable_queryCmd.Flags().Bool("recurse", false, "Recurse into subfragments")
	repository_store_immutableCmd.AddCommand(repository_store_immutable_queryCmd)

	carapace.Gen(repository_store_immutable_queryCmd).PositionalCompletion(
		carapace.ActionValues(), // address
	)
}
