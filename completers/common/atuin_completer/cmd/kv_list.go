package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var kv_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all keys in a namespace, or in all namespaces",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_listCmd).Standalone()

	kv_listCmd.Flags().Bool("all", false, "List all keys in all namespaces")
	kv_listCmd.Flags().BoolP("all-namespaces", "a", false, "List all keys in all namespaces")
	kv_listCmd.Flags().BoolP("help", "h", false, "Print help")
	kv_listCmd.Flags().StringP("namespace", "n", "", "Namespace to list keys from")
	kv_listCmd.Flag("all").Hidden = true
	kvCmd.AddCommand(kv_listCmd)

	carapace.Gen(kv_listCmd).FlagCompletion(carapace.ActionMap{
		"namespace": atuin.ActionNamespaces(),
	})
}
