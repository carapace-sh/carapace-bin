package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update dependencies as recorded in the local lock file",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("update"),
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("aggressive", false, "Force updating all dependencies of SPEC as well when used with -p")
	updateCmd.Flags().Bool("dry-run", false, "Don't actually write the lockfile")
	updateCmd.Flags().BoolP("help", "h", false, "Print help")
	updateCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	updateCmd.Flags().StringSliceP("package", "p", []string{}, "Package to update")
	updateCmd.Flags().String("precise", "", "Update a single dependency to exactly PRECISE when used with -p")
	updateCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	updateCmd.Flags().BoolP("workspace", "w", false, "Only update the workspace packages")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(updateCmd, false),
	})
}
