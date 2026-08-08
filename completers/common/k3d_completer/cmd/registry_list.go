package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var registry_listCmd = &cobra.Command{
	Use:     "list [NAME [NAME...]]",
	Short:   "List registries",
	Aliases: []string{"ls", "get"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_listCmd).Standalone()

	registry_listCmd.Flags().Bool("no-headers", false, "Disable headers")
	registry_listCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml")
	registryCmd.AddCommand(registry_listCmd)

	carapace.Gen(registry_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("json", "yaml"),
	})

	carapace.Gen(registry_listCmd).PositionalAnyCompletion(
		k3d.ActionRegistries().FilterArgs(),
	)
}
