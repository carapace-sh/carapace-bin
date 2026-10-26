package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:     "depends",
	Short:   "Shows which packages cause the given package to be installed",
	Aliases: []string{"why"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()

	dependsCmd.Flags().Bool("locked", false, "Read dependency information from composer.lock")
	dependsCmd.Flags().BoolP("recursive", "r", false, "Recursively resolves up to the root package")
	dependsCmd.Flags().BoolP("tree", "t", false, "Prints the results as a nested tree")

	rootCmd.AddCommand(dependsCmd)

	carapace.Gen(dependsCmd).PositionalCompletion(
		action.ActionInstalledPackages(dependsCmd),
	)
}
