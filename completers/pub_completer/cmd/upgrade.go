package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pub_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the current package's dependencies to latest versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change any.")
	upgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	upgradeCmd.Flags().Bool("major-versions", false, "Upgrades packages to their latest resolvable versions")
	upgradeCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	upgradeCmd.Flags().Bool("no-precompile", false, "Do not precompile executables in immediate dependencies.")
	upgradeCmd.Flags().Bool("null-safety", false, "Upgrade constraints in pubspec.yaml to null-safety versions")
	upgradeCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	upgradeCmd.Flags().Bool("precompile", false, "Precompile executables in immediate dependencies.")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionDependencies().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
