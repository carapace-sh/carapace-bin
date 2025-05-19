package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependency_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the dependencies for the given chart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependency_listCmd).Standalone()
	dependency_listCmd.Flags().Uint("max-col-width", 80, "maximum column width for output table")
	dependencyCmd.AddCommand(dependency_listCmd)

	carapace.Gen(dependency_listCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
