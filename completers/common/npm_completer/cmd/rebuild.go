package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var rebuildCmd = &cobra.Command{
	Use:     "rebuild",
	Short:   "Rebuild a package",
	Aliases: []string{"rb"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildCmd).Standalone()
	rebuildCmd.Flags().StringArray("allow-scripts", nil, "packages whose install-time lifecycle scripts are allowed to run")
	rebuildCmd.Flags().Bool("bin-links", false, "create symlinks for package executables")
	rebuildCmd.Flags().Bool("dangerously-allow-all-scripts", false, "bypass the allowScripts policy entirely")
	rebuildCmd.Flags().Bool("foreground-scripts", false, "run build scripts in the foreground process")
	rebuildCmd.Flags().BoolP("global", "g", false, "operate globally")
	rebuildCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	rebuildCmd.Flags().Bool("no-bin-links", false, "skip symlinks for package executables")
	rebuildCmd.Flags().Bool("strict-allow-scripts", false, "turn install-script policy from warning into error")
	addWorkspaceFlags(rebuildCmd)

	rootCmd.AddCommand(rebuildCmd)

	carapace.Gen(rebuildCmd).PositionalAnyCompletion(
		npm.ActionDependencyNames(),
	)
}
