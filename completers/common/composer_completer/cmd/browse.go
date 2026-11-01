package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:     "browse",
	Short:   "Opens the package's repository URL or homepage in your browser",
	Aliases: []string{"home"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(browseCmd).Standalone()

	browseCmd.Flags().BoolP("homepage", "H", false, "Open the homepage instead of the repository URL.")
	browseCmd.Flags().BoolP("show", "s", false, "Only show the homepage or repository URL.")

	rootCmd.AddCommand(browseCmd)

	carapace.Gen(browseCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(browseCmd).FilterArgs(),
	)
}
