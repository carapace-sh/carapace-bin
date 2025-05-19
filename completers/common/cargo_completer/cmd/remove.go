package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove dependencies from a Cargo.toml manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("build", false, "Remove from build-dependencies")
	removeCmd.Flags().Bool("dev", false, "Remove from dev-dependencies")
	removeCmd.Flags().BoolP("dry-run", "n", false, "Don't actually write the manifest")
	removeCmd.Flags().BoolP("help", "h", false, "Print help")
	removeCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	removeCmd.Flags().StringP("package", "p", "", "Package to remove from")
	removeCmd.Flags().String("target", "", "Remove from target-dependencies")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(removeCmd, false),
		"target":        action.ActionTargets(removeCmd, action.TargetOpts{}.Default()),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		action.ActionDependencies(removeCmd, false).FilterArgs(),
	)
}
