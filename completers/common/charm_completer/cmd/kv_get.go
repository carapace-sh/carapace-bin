package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var kv_getCmd = &cobra.Command{
	Use:   "get KEY[@DB]",
	Short: "Get a value for a key with an optional @ db.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_getCmd).Standalone()

	kv_getCmd.Flags().BoolP("show-binary", "b", false, "print binary values")
	kvCmd.AddCommand(kv_getCmd)

	carapace.Gen(kv_getCmd).PositionalCompletion(
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
