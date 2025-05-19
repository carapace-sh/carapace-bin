package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()
	uninstallCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	uninstallCmd.Flags().BoolP("save", "S", false, "Package will appear in your `dependencies`")
	uninstallCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	uninstallCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	uninstallCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	uninstallCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	addWorkspaceFlags(uninstallCmd)

	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		npm.ActionModules(),
	)
}
