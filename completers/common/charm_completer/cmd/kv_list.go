package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var kv_listCmd = &cobra.Command{
	Use:   "list [@DB]",
	Short: "List all key value pairs with an optional @ db.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_listCmd).Standalone()

	kv_listCmd.Flags().StringP("delimiter", "d", "", "delimiter to separate keys and values")
	kv_listCmd.Flags().BoolP("keys-only", "k", false, "only print keys and don't fetch values from the db")
	kv_listCmd.Flags().BoolP("reverse", "r", false, "list in reverse lexicographic order")
	kv_listCmd.Flags().BoolP("show-binary", "b", false, "print binary values")
	kv_listCmd.Flags().BoolP("values-only", "v", false, "only print values")
	kvCmd.AddCommand(kv_listCmd)

	carapace.Gen(kv_listCmd).PositionalCompletion(
		charm.ActionDatabases(),
	)
}
