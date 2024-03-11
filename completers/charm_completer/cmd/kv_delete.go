package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var kv_deleteCmd = &cobra.Command{
	Use:   "delete KEY[@DB]",
	Short: "Delete a key with an optional @ db.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_deleteCmd).Standalone()

	kvCmd.AddCommand(kv_deleteCmd)

	carapace.Gen(kv_deleteCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return charm.ActionDatabases()
			}
		}),
	)
}
