package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var kv_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete one or more key-value pairs",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_deleteCmd).Standalone()

	kv_deleteCmd.Flags().BoolP("help", "h", false, "Print help")
	kv_deleteCmd.Flags().StringP("namespace", "n", "", "Namespace for the key-value pair")
	kvCmd.AddCommand(kv_deleteCmd)

	carapace.Gen(kv_deleteCmd).FlagCompletion(carapace.ActionMap{
		"namespace": atuin.ActionNamespaces(),
	})

	carapace.Gen(kv_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return atuin.ActionKeys(kv_deleteCmd.Flag("namespace").Value.String())
		}),
	)
}
