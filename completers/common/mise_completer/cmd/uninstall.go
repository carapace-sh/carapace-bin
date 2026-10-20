package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove installed tool versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().BoolP("all", "a", false, "Remove every installed version of the named tools, or of every tool when none is named")
	uninstallCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	uninstallCmd.Flags().Bool("dry-run-code", false, "Like --dry-run, but exit with code 1 if there are tools to uninstall")
	uninstallCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		mise.ActionInstalledToolVersions(),
	)
}
