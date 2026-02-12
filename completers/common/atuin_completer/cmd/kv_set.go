package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var kv_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a key-value pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_setCmd).Standalone()

	kv_setCmd.Flags().BoolP("help", "h", false, "Print help")
	kv_setCmd.Flags().StringP("key", "k", "", "Key to set")
	kv_setCmd.Flags().StringP("namespace", "n", "", "Namespace for the key-value pair")
	kv_setCmd.MarkFlagRequired("key")
	kvCmd.AddCommand(kv_setCmd)

	carapace.Gen(kv_setCmd).FlagCompletion(carapace.ActionMap{
		"namespace": atuin.ActionNamespaces(),
	})

	carapace.Gen(kv_setCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return atuin.ActionKeys(kv_setCmd.Flag("namespace").Value.String())
		}),
	)
}
