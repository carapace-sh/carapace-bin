package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pub_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a dependency from the current package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("[no-]offline", false, "Use cached packages instead of accessing the network.")
	removeCmd.Flags().Bool("[no-]precompile", false, "Precompile executables in immediate dependencies.")
	removeCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change")
	removeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalCompletion(
		action.ActionDependencies(),
	)
}
