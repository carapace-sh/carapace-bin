package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var kv_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a saved value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_getCmd).Standalone()

	kv_getCmd.Flags().BoolP("help", "h", false, "Print help")
	kv_getCmd.Flags().StringP("namespace", "n", "", "Namespace for the key-value pair")
	kvCmd.AddCommand(kv_getCmd)

	carapace.Gen(kv_getCmd).FlagCompletion(carapace.ActionMap{
		"namespace": atuin.ActionNamespaces(),
	})

	carapace.Gen(kv_getCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return atuin.ActionKeys(kv_getCmd.Flag("namespace").Value.String())
		}),
	)
}
