package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var prohibitsCmd = &cobra.Command{
	Use:     "prohibits",
	Short:   "Shows which packages prevent the given package from being installed",
	Aliases: []string{"why-not"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prohibitsCmd).Standalone()

	prohibitsCmd.Flags().Bool("locked", false, "Read dependency information from composer.lock")
	prohibitsCmd.Flags().BoolP("recursive", "r", false, "Recursively resolves up to the root package")
	prohibitsCmd.Flags().BoolP("tree", "t", false, "Prints the results as a nested tree")

	rootCmd.AddCommand(prohibitsCmd)

	carapace.Gen(prohibitsCmd).PositionalCompletion(
		action.ActionInstalledPackages(prohibitsCmd),
	)
}
