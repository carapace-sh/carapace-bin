package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var namespace_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List OpenFaaS namespaces",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespace_listCmd).Standalone()

	namespace_listCmd.Flags().BoolP("json", "j", false, "Output namespaces as JSON")
	namespaceCmd.AddCommand(namespace_listCmd)
}
