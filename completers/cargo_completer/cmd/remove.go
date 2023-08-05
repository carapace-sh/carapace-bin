package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove dependencies from a Cargo.toml manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("build", false, "Remove as build dependency")
	removeCmd.Flags().Bool("dev", false, "Remove as development dependency")
	removeCmd.Flags().Bool("dry-run", false, "Don't actually write the manifest")
	removeCmd.Flags().BoolP("help", "h", false, "Print help")
	removeCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	removeCmd.Flags().StringP("package", "p", "", "Package to remove from")
	removeCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	removeCmd.Flags().String("target", "", "Remove as dependency from the given target platform")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(removeCmd, false),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		action.ActionDependencies(removeCmd, false).FilterArgs(),
	)
}
