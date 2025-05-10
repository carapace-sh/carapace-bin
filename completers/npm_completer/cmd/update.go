package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	updateCmd.Flags().Bool("audit", false, "submit audit reports")
	updateCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	updateCmd.Flags().Bool("dry-run", false, "Only report changes")
	updateCmd.Flags().Bool("fund", false, "Display funding message")
	updateCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	updateCmd.Flags().Bool("global-style", false, "Use global layout")
	updateCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	updateCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	updateCmd.Flags().StringArray("omit", nil, "omit dependency types")
	updateCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	updateCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(updateCmd)

	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		npm.ActionModules(), // TODO support global
	)
}
