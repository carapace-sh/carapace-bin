package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "View skaffold namespace information for resources it manages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_namespacesCmd).Standalone()

	applyCommonFlagCompletion(inspect_namespacesCmd)
	inspectCmd.AddCommand(inspect_namespacesCmd)
}
