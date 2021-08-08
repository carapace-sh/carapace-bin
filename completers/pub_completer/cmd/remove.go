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

	removeCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change")
	removeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	removeCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	removeCmd.Flags().Bool("no-precompile", false, "Do not precompile executables in immediate dependencies.")
	removeCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	removeCmd.Flags().Bool("precompile", false, "Precompile executables in immediate dependencies.")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalCompletion(
		action.ActionDependencies(),
	)
}
