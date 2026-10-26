package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var suggestsCmd = &cobra.Command{
	Use:   "suggests",
	Short: "Shows package suggestions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suggestsCmd).Standalone()

	suggestsCmd.Flags().BoolP("all", "a", false, "Show suggestions from all dependencies, including transitive ones")
	suggestsCmd.Flags().Bool("by-package", false, "Groups output by suggesting package (default)")
	suggestsCmd.Flags().Bool("by-suggestion", false, "Groups output by suggested package")
	suggestsCmd.Flags().Bool("list", false, "Show only list of suggested package names")
	suggestsCmd.Flags().Bool("no-dev", false, "Exclude suggestions from require-dev packages")

	rootCmd.AddCommand(suggestsCmd)

	carapace.Gen(suggestsCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(suggestsCmd).FilterArgs(),
	)
}
