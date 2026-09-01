package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:     "view",
	Short:   "View registry information about a package",
	Aliases: []string{"info", "show", "v"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewCmd).Standalone()

	viewCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	viewCmd.Flags().Bool("json", false, "Show information in JSON format")
	viewCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(viewCmd)

	carapace.Gen(viewCmd).PositionalCompletion(
		npm.ActionPackageSearch(""),
	)
}
