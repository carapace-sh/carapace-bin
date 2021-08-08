package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the current package's dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change any.")
	getCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	getCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	getCmd.Flags().Bool("no-precompile", false, "Do not precompile executables in immediate dependencies.")
	getCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	getCmd.Flags().Bool("precompile", false, "Precompile executables in immediate dependencies.")
	rootCmd.AddCommand(getCmd)
}
