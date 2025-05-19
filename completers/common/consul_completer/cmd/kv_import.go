package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var kv_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a tree stored as JSON to the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_importCmd).Standalone()
	addClientFlags(kv_importCmd)
	addServerFlags(kv_importCmd)
	kv_importCmd.Flags().String("namespace", "", "Specifies the namespace to query.")

	kvCmd.AddCommand(kv_importCmd)

	// TODO namespace completion

	carapace.Gen(kv_importCmd).PositionalCompletion(
		action.ActionOptionalFiles(".json"),
	)
}
