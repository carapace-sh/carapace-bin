package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var kv_setCmd = &cobra.Command{
	Use:   "set KEY[@DB] VALUE",
	Short: "Set a value for a key with an optional @ db.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_setCmd).Standalone()

	kvCmd.AddCommand(kv_setCmd)

	carapace.Gen(kv_setCmd).PositionalCompletion(
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
