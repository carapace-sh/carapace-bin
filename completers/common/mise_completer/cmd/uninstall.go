package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Removes installed tool versions",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().BoolP("all", "a", false, "Remove all installed tools")
	uninstallCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed without actually removing")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		action.ActionInstalledToolVersions(),
	)
}
