package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/asdf"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove a specific version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalCompletion(
		asdf.ActionInstalledPlugins(),
		asdf.ActionInstalledVersions(""),
	)
}
