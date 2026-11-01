package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var namespace_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespace_createCmd).Standalone()

	namespace_createCmd.Flags().StringArray("annotation", nil, "Set one or more annotation (ANNOTATION=VALUE)")
	namespace_createCmd.Flags().StringArrayP("label", "l", nil, "Set one or more label (LABEL=VALUE)")
	namespaceCmd.AddCommand(namespace_createCmd)
}
